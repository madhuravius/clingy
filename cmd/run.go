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

		if err := lib.ClingyCanRun(); err != nil {
			logger.Println("Unable to run clingy due to error in startup", err)
			fmt.Println(fmt.Sprintf("Clingy cannot run for reason: %s", err.Error()))
			os.Exit(1)
		}

		clingyData, err := lib.ParseClingyFile(logger, inputFile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error in reading: %s", inputFile), err)
			os.Exit(1)
		}
		fmt.Printf("Running: %s", clingyData.Label)

		for idx, step := range clingyData.Steps {
			// clear terminal for fresh screenshot
			internal.ClearTerminal()
			args := step.Args
			if args == nil {
				args = []string{}
			}

			// execute actual command
			_, err := internal.ExecuteCommand(logger, step.Command, args)
			if err != nil {
				logger.Println("Error in executing command", err)
				os.Exit(1)
			}

			// capture window before proceeding
			imagePath, err := internal.CaptureWindow(logger, getOutputPath(), strconv.Itoa(idx), ".jpg")
			if err != nil {
				logger.Println("Error in capturing image", err)
				os.Exit(1)
			}

			// if image-only report, add labels/descriptions to the image itself
			if reportStyle == "images-only" {
				if step.Label != "" {
					if err := internal.AddLabelToImage(logger, step.Label, imagePath); err != nil {
						logger.Println("Error in adding label to image", err)
						os.Exit(1)
					}
				}
				if step.Description != "" {
					if err := internal.AddDescriptionToImage(logger, step.Description, imagePath); err != nil {
						logger.Println("Error in adding description to image", err)
						os.Exit(1)
					}
				}
			}
			clingyData.Steps[idx].ImageOutput = fmt.Sprintf("%s%s", strconv.Itoa(idx), ".jpg")
		}

		internal.ClearTerminal()
		switch reportStyle {
		case "images-only":
			fmt.Println(fmt.Sprintf("Completed clingy run, generated images at %s.", getOutputPath()))
		case "html-simple":
			fmt.Println("Completed clingy run, generating report.")

			reportPath := fmt.Sprintf("%s/index.html", getOutputPath())
			if err := internal.GenerateHTMLReport(logger, clingyData, reportPath); err != nil {
				fmt.Println("Error in generating report")
				os.Exit(1)
			}
			fmt.Println(fmt.Sprintf("Generated report: %s", reportPath))
		}
	},
}
