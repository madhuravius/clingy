package lib

import (
	"errors"
	"regexp"
)

type MatchingType string

const (
	// Regex - basic regex matching. See regexp.Compile - https://pkg.go.dev/regexp#Compile
	Regex MatchingType = "regex"
	// Positional - based on a delimiter and position. See - strings.Split - https://pkg.go.dev/strings#Split
	Positional MatchingType = "positional"
	// Full - fully captured as input
	Full MatchingType = "full"
)

var (
	// NoMatchError - when no match is found in output args, but one is expected
	NoMatchError = errors.New("expected match from output, but no match in output found from matching_args")
)

var (
	// RegexMatchingInput - a regex specifically to target inputs (see HydrateArgs)
	RegexMatchingInput = regexp.MustCompile(`\$\[\[([^]]+)\]`)
)

// MatchingArgs - specific wildcard matches for an assorted set of possible uses
type MatchingArgs struct {
	// Regex-related values
	// Regex - raw regex string to use in regexp library
	Regex string `yaml:"regex"`

	// Positional-related values
	// PositionalDelimiter - a value to delimit a string on
	PositionalDelimiter string `yaml:"positional_delimiter"`
	// PositionalIndex - a value that represents an index in a search
	PositionalIndex int `yaml:"positional_index"`
}

// ClingyOutputProcessing - structure that tells clingy what to do with an output
type ClingyOutputProcessing struct {
	Key          string       `yaml:"key"`
	MatchingType MatchingType `yaml:"matching_type"`
	// MatchingArgs - for a given match type to extract output, args will be needed to check for matches
	MatchingArgs MatchingArgs `yaml:"matching_args"`
	// FailOnNoMatch - forces a failure on a step if no match found (by default false)
	FailOnNoMatch bool `yaml:"fail_on_no_match"`
}

// ClingyStep - step to execute on
type ClingyStep struct {
	Label            string                  `yaml:"label"`
	Description      string                  `yaml:"description"`
	Command          string                  `yaml:"command"`
	Args             []string                `yaml:"args"`
	OutputProcessing *ClingyOutputProcessing `yaml:"output_processing"`
	ImageOutput      string
}

// ClingyTemplate - a full set of clingy instruction to follow
type ClingyTemplate struct {
	Label       string       `yaml:"label"`
	Description string       `yaml:"description"`
	Steps       []ClingyStep `yaml:"steps"`
	// StepOutputs - for each step key (see: ClingyOutputProcessing), a corresponding output can be stored
	// and this output can be reused in execution by reference via string substitution
	StepOutputs map[string]string
}
