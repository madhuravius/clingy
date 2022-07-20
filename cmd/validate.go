package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"clingy/lib"
)

// newValidateCmd - Command to validate the yaml
func (r *RootConfig) newValidateCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "validate",
		Short:  "Validate a clingy.yml file",
		PreRun: initRunWithoutArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Validating %s\n", inputFile)

			if err := lib.ClingyCanRun(); err != nil {
				fmt.Printf("Error in checking if clingy can run: %s\n", err.Error())
				r.ExitTools.Exit(1)
			}

			_, err := lib.ParseClingyFile(logger, inputFile)
			if err != nil {
				fmt.Printf("Error in validating: %s %s", inputFile, err.Error())
				r.ExitTools.Exit(1)
			}

			cmd.Println("Completed validation, looks good!")
			r.ExitTools.Exit(0)
		},
	}
}
