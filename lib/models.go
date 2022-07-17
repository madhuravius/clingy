package lib

// ClingyStep - step to execute on
type ClingyStep struct {
	Label       string `yaml:"label"`
	Description string `yaml:"description"`
	Command     string `yaml:"command"`
}

// ClingyTemplate - a full set of clingy instruction to follow
type ClingyTemplate struct {
	Label       string       `yaml:"label"`
	Description string       `yaml:"description"`
	Steps       []ClingyStep `yaml:"steps"`
}
