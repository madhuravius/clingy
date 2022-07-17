package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"clingy/internal"
	"clingy/lib"
)

// validateCmd - Command to validate the yaml
var validateCmd = &cobra.Command{
	Use:    "validate",
	Short:  "Validate a clingy.yml file",
	PreRun: initRunWithoutArtifactDirectoryCreate,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Validating %s\n", inputFile)

		if err := internal.CheckMagickBinary(); err != nil {
			fmt.Println("Error with magick binary", err)
			os.Exit(1)
		}

		_, err := lib.ParseClingyFile(inputFile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error in validating: %s", inputFile), err)
			os.Exit(1)
		}

		fmt.Println("Completed validation, looks good!")
		os.Exit(0)
	},
}
