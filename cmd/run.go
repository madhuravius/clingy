package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"clingy/internal"
	"clingy/lib"
)

// runCmd - actually run clingy
var runCmd = &cobra.Command{
	Use:    "run",
	Short:  "Run clingy",
	PreRun: initRunWithArtifactDirectoryCreate,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Running clingy")

		clingyData, err := lib.ParseClingyFile(logger, inputFile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error in reading: %s", inputFile), err)
			os.Exit(1)
		}
		fmt.Printf("Running: %s", clingyData.Label)

		for idx, step := range clingyData.Steps {
			internal.ClearTerminal()
			args := step.Args
			if args == nil {
				args = []string{}
			}
			imagePath := internal.CaptureWindow(logger, getOutputPath(), step.Command, args, strconv.Itoa(idx))
			if step.Label != "" {
				internal.AddLabelToImage(logger, step.Label, imagePath)
			}
			if step.Description != "" {
				internal.AddDescriptionToImage(logger, step.Description, imagePath)
			}
		}

		internal.ClearTerminal()
		fmt.Println("Completed clingy run.")
	},
}
