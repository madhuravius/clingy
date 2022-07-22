package lib

import (
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

// GetMagickBinary - check if imagemagick binary found in path
func GetMagickBinary() (string, error) {
	magickPath := "magick"
	if _, magickErr := exec.LookPath(magickPath); magickErr != nil {
		// fall back to import and see if that exists
		magickPath = "import"
		if _, err := exec.LookPath(magickPath); err != nil {
			return "", errors.New("error: magick or fallback binary (import) not found")
		}
	}
	return magickPath, nil
}

func CheckScreenshotBinary() error {
	if _, err := exec.LookPath("screenshot"); err != nil {
		return errors.New("error: screenshot python3 bin not found")
	}
	return nil
}

// ClingyCanRun - catch-all for ensuring clingy can actually run
func ClingyCanRun() error {
	if _, err := GetMagickBinary(); err != nil {
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

	// hydrate the yaml data from environment variables in parent context
	data = []byte(os.ExpandEnv(string(data)))

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
