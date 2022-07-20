package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"clingy/internal"
	"clingy/lib"
)

// newRunCmd - actually run clingy
func (r *RootConfig) newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "run",
		Short:  "Run clingy",
		PreRun: initRunWithArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Running clingy")

			if err := lib.ClingyCanRun(); err != nil {
				logger.Println("Unable to run clingy due to error in startup", err)
				fmt.Println(fmt.Sprintf("Clingy cannot run for reason: %s", err.Error()))
				r.ExitTools.Exit(1)
			}

			clingyData, err := lib.ParseClingyFile(logger, inputFile)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error in reading: %s", inputFile), err)
				r.ExitTools.Exit(1)
			}
			fmt.Printf("Running: %s", clingyData.Label)

			for idx, step := range clingyData.Steps {
				// clear terminal for fresh screenshot
				internal.ClearTerminal()
				if step.Args == nil {
					step.Args = []string{}
				} else {
					// preprocess args for input via string substitution from prior outputs
					if err := lib.HydrateArgs(logger, clingyData, idx); err != nil {
						fmt.Println("Error in using output map in input arguments", err)
						r.ExitTools.Exit(1)
					}
				}

				// execute actual command
				output, err := internal.ExecuteCommand(logger, step.Command, step.Args)
				if err != nil {
					logger.Println("Error in executing command", err)
					r.ExitTools.Exit(1)
				}

				// capture window before proceeding
				imagePath, err := r.Magick.CaptureWindow(logger, getOutputPath(), strconv.Itoa(idx), ".jpg")
				if err != nil {
					logger.Println("Error in capturing image", err)
					r.ExitTools.Exit(1)
				}
				clingyData.Steps[idx].ImageOutput = fmt.Sprintf("%s%s", strconv.Itoa(idx), ".jpg")

				// if image-only report, add labels/descriptions to the image itself
				if reportStyle == "images-only" {
					if step.Label != "" {
						if err := r.Magick.AddLabelToImage(logger, step.Label, imagePath); err != nil {
							logger.Println("Error in adding label to image", err)
							r.ExitTools.Exit(1)
						}
					}
					if step.Description != "" {
						if err := r.Magick.AddDescriptionToImage(logger, step.Description, imagePath); err != nil {
							logger.Println("Error in adding description to image", err)
							r.ExitTools.Exit(1)
						}
					}
				}

				// if output key, process output and store it for future use
				if step.OutputProcessing != nil {
					fmt.Println(fmt.Sprintf("Output processing found for key %s", step.OutputProcessing.Key))
					if err := lib.HydrateOutput(logger, string(output), clingyData, idx); err != nil {
						fmt.Println("Error in capturing output in processing", err)
						r.ExitTools.Exit(1)
					}
					logger.Println(fmt.Sprintf("Finished processing output: %s", output))
				}
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
					r.ExitTools.Exit(1)
				}
				fmt.Println(fmt.Sprintf("Generated report: %s", reportPath))
			}
		},
	}

}
