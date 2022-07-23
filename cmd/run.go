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
				cmd.Printf("Clingy cannot run for reason: %s\n", err.Error())
				r.ExitTools.Exit(1)
			}

			clingyData, err := lib.ParseClingyFile(logger, inputFile)
			if err != nil {
				cmd.Printf("Error in reading: %s %s\n", inputFile, err.Error())
				r.ExitTools.Exit(1)
			}
			cmd.Printf("Running: %s", clingyData.Label)

			for idx, step := range clingyData.Steps {
				// clear terminal for fresh screenshot
				internal.ClearTerminal()
				if step.Args == nil {
					step.Args = []string{}
				} else {
					// preprocess args for input via string substitution from prior outputs
					if err := lib.HydrateArgs(logger, clingyData, idx); err != nil {
						cmd.Println("Error in using output map in input arguments", err)
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
				imagePath, err := r.ImageTools.CaptureWindow(logger, getOutputPath(), strconv.Itoa(idx), ".jpg")
				if err != nil {
					logger.Println("Error in capturing image", err)
					r.ExitTools.Exit(1)
				}
				clingyData.Steps[idx].ImageOutput = fmt.Sprintf("%s%s", strconv.Itoa(idx), ".jpg")

				// if image-only report, add labels/descriptions to the image itself
				if reportStyle == "images-only" {
					if step.Label != "" {
						if err := r.ImageTools.AddLabelToImage(logger, imagePath, step.Label); err != nil {
							logger.Println("Error in adding label to image", err)
							r.ExitTools.Exit(1)
						}
					}
					if step.Description != "" {
						if err := r.ImageTools.AddDescriptionToImage(logger, imagePath, step.Description); err != nil {
							logger.Println("Error in adding description to image", err)
							r.ExitTools.Exit(1)
						}
					}
				}

				// if output key, process output and store it for future use
				if step.OutputProcessing != nil {
					cmd.Printf("Output processing found for key %s\n", step.OutputProcessing.Key)
					if err := lib.HydrateOutput(logger, string(output), clingyData, idx); err != nil {
						cmd.Println("Error in capturing output in processing", err)
						r.ExitTools.Exit(1)
					}
					logger.Printf("Finished processing output: %s", output)
				}
			}

			internal.ClearTerminal()
			switch reportStyle {
			case "images-only":
				cmd.Printf("Completed clingy run, generated images at %s.\n", getOutputPath())
			case "carousel", "html-simple":
				cmd.Println("Completed clingy run, generating report.")

				reportPath := fmt.Sprintf("%s/index.html", getOutputPath())
				if err := internal.GenerateHTMLReport(logger, clingyData, reportStyle, reportPath); err != nil {
					cmd.Println("Error in generating report")
					r.ExitTools.Exit(1)
				}
				cmd.Printf("Generated report: %s\n", reportPath)
			}
		},
	}
}
