package cmd

import (
	"fmt"
	"os"

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
				fmt.Println(fmt.Sprintf("Error in checking if clingy can run: %s", err.Error()))
				os.Exit(1)
			}

			_, err := lib.ParseClingyFile(logger, inputFile)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error in validating: %s", inputFile), err)
				os.Exit(1)
			}

			fmt.Println("Completed validation, looks good!")
			os.Exit(0)
		},
	}
}
