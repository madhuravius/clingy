package lib

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

var (
	NoStepsError = errors.New("error: unable to process template, no steps")
)

// CheckMagickBinary - check if imagemagick binary found in path
func CheckMagickBinary() error {
	if _, err := exec.LookPath("magick"); os.IsNotExist(err) {
		return errors.New("error: magick binary not found")
	}
	return nil
}

// ClingyCanRun - catch-all for ensuring clingy can actually run
func ClingyCanRun() error {
	if err := CheckMagickBinary(); err != nil {
		return err
	}

	if os.Getenv("WINDOWID") == "" {
		return errors.New("environment variable WINDOWID required to proceed")
	}

	return nil
}

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
