package lib

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	NoStepsError = errors.New("error: unable to process template, no steps")
)

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
