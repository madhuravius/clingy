package lib

import (
	"clingy/internal/images"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	// NoStepsError ...
	NoStepsError = errors.New("error: unable to process template, no steps")
)

// CheckMagickBinary - check if imagemagick binary found in path
func CheckMagickBinary() error {
	if _, err := exec.LookPath("magick"); os.IsNotExist(err) {
		return errors.New("error: magick binary not found")
	}
	return nil
}

func CheckScreenshotBinary() error {
	if _, err := exec.LookPath("screenshot"); os.IsNotExist(err) {
		return errors.New("error: screenshot python3 bin not found")
	}
	return nil
}

func GetClingyImageCapture() images.ImageProcessingImpl {
	switch runtime.GOOS {
	case "darwin":
		return images.NewMacScreenshotClient()
	case "linux":
		return images.NewMagickClient()
	default:
		fmt.Println("WARNING - Operating system not supported, trying to use imagemagick client.")
		return images.NewMagickClient()
	}
}

// ClingyCanRun - catch-all for ensuring clingy can actually run
func ClingyCanRun() error {
	if err := CheckMagickBinary(); err != nil {
		return err
	}

	switch runtime.GOOS {
	case "darwin":
		if err := CheckScreenshotBinary(); err != nil {
			return err
		}
		if os.Getenv("WINDOW_NAME") == "" {
			return errors.New("environment variable WINDOW_NAME required to proceed")
		}
	case "linux":
		if os.Getenv("WINDOWID") == "" {
			return errors.New("environment variable WINDOWID required to proceed")
		}
	default:
		fmt.Println("Warning, running on an unsupported OS")
		if os.Getenv("WINDOWID") == "" {
			return errors.New("environment variable WINDOWID required to proceed")
		}
	}

	return nil
}

// ParseClingyFile - unmarshal a target filename to a clingy template pointer for reuse
func ParseClingyFile(logger *log.Logger, fileName string) (*ClingyTemplate, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var clingyData ClingyTemplate
	if err = yaml.Unmarshal(data, &clingyData); err != nil {
		return nil, err
	}

	if len(clingyData.Steps) == 0 {
		logger.Println("Error - missing steps to execute in order")
		return nil, NoStepsError
	}

	return &clingyData, nil
}
