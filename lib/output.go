package lib

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// HydrateArgs - change args by reference with interpolated values expected from map, optionally error
func HydrateArgs(logger *log.Logger, clingyData *ClingyTemplate, step int) error {
	logger.Println(fmt.Sprintf("Hydrating output (step %d) and args", step), clingyData.Steps[step].Args)
	for argIdx, arg := range clingyData.Steps[step].Args {
		if RegexMatchingInput.MatchString(arg) {
			logger.Println("Arg match found", arg)
			rawOutputKeyToMatch := RegexMatchingInput.FindStringSubmatch(arg)
			if rawOutputKeyToMatch == nil {
				return errors.New(fmt.Sprintf(
					"in hydrating inputs for step, output anticipated key %s, but not found in outputs!",
					rawOutputKeyToMatch))
			}
			if output, ok := clingyData.StepOutputs[rawOutputKeyToMatch[1]]; !ok {
				return errors.New(fmt.Sprintf(
					"in hydrating inputs for step, output anticipated key %s, but not found in outputs!",
					rawOutputKeyToMatch[1]))
			} else {
				clingyData.Steps[step].Args[argIdx] = output
			}
		}
	}
	logger.Println(fmt.Sprintf("Hydrated output (step %d)", step))
	return nil
}

// HydrateOutput - store output for future consumption
func HydrateOutput(logger *log.Logger, output string, clingyData *ClingyTemplate, step int) error {
	outputArgs := clingyData.Steps[step].OutputProcessing
	logger.Println(fmt.Sprintf(
		"Starting to hydrate output (step %d) for %s's clingyData on key: %s",
		step,
		clingyData.Label,
		outputArgs.Key,
	))

	var value string
	switch outputArgs.MatchingType {
	case Positional:
		rawOutput := strings.Split(output, outputArgs.MatchingArgs.PositionalDelimiter)
		if len(rawOutput) >= outputArgs.MatchingArgs.PositionalIndex && outputArgs.FailOnNoMatch {
			return NoMatchError
		}
		value = rawOutput[outputArgs.MatchingArgs.PositionalIndex]
	case Regex:
		regex, err := regexp.Compile(outputArgs.MatchingArgs.Regex)
		if err != nil {
			return err
		}
		rawValue := regex.FindStringSubmatch(output)
		if (len(rawValue) == 0 || rawValue[1] == "") && outputArgs.FailOnNoMatch {
			return NoMatchError
		} else if len(rawValue) >= 1 {
			value = rawValue[1]
		}
	case Full:
		value = output
	default:
		return errors.New("in populating outputs for future inputs, no acceptable output processing matching type provided")
	}

	if clingyData.StepOutputs == nil {
		clingyData.StepOutputs = make(map[string]string)
	}
	clingyData.StepOutputs[outputArgs.Key] = value
	logger.Println("Hydrated output", clingyData.StepOutputs)

	return nil
}
