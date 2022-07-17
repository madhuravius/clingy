package lib

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ParseClingyFile(fileName string) (*ClingyTemplate, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var clingyData ClingyTemplate
	if err = yaml.Unmarshal(data, &clingyData); err != nil {
		return nil, err
	}

	return &clingyData, nil
}
