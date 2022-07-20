package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"clingy/internal"
	"github.com/spf13/cobra"
)

var (
	// buildNumber - a way to distinguish between various builds within an output directory
	buildNumber = strconv.Itoa(int(time.Now().Unix()))
	// logger - logger for debugging reasons, init'ed and typically writes to file in output directory w/ build #
	logger *log.Logger
	// version - version of the app to spit out, currently manually set :(
	version = "v0.3.0"

	// flags
	// debug - enable verbose logging
	debug bool
	// outputPath - a location to dump artifacts/output as a result of a clingy run
	outputPath = "./output"
	// inputFile - a path that contains an input file to digest and run clingy against
	inputFile = "./.clingy.yaml"
	// reportStyle - output format to share screenshots
	reportStyle = "html-simple"
)

// RootConfig - variables to pass in for reuse and testing
type RootConfig struct {
	Magick internal.MagickClientImpl
}

// getOutputPath - a string that generates a union of an (dynamic) output path and build number for artifacts
func getOutputPath() string {
	return fmt.Sprintf("%s/%s", outputPath, buildNumber)
}

// initRunWithArtifactDirectoryCreate - use this when needing to create a destination directory (ex: `run`)
func initRunWithArtifactDirectoryCreate(_ *cobra.Command, _ []string) {
	internal.InitDestinationDirectory(getOutputPath())
	logger = internal.InitLogger(getOutputPath(), debug)
}

// initRunWithoutArtifactDirectoryCreate - use this function when not needing a generalized create except with debug
func initRunWithoutArtifactDirectoryCreate(_ *cobra.Command, _ []string) {
	if debug {
		internal.InitDestinationDirectory(getOutputPath())
	}
	logger = internal.InitLogger(getOutputPath(), debug)
}

// RootCmd - entrypoint for clingy app
func RootCmd(c *RootConfig) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:    "clingy",
		Short:  "clingy is a tool to test and capture CLI flows",
		Long:   `clingy is a tool to test and capture CLI flows.`,
		PreRun: initRunWithoutArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				if err := cmd.Help(); err != nil {
					logger.Println("Error when trying to print help.", err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		},
	}

	rootCmd.AddCommand(c.newCleanCmd())
	rootCmd.AddCommand(c.newInitCmd())
	rootCmd.AddCommand(c.newRunCmd())
	rootCmd.AddCommand(c.newValidateCmd())
	rootCmd.AddCommand(c.newVersionCmd())

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logs")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "outputPath", "o", outputPath, "build path that dumps outputs")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "inputFile", "i", inputFile, "input file representing a .clingy.yaml")
	rootCmd.PersistentFlags().StringVarP(&reportStyle, "reportStyle", "r", reportStyle, "report style to output to (choices: 'html-simple', 'images-only')")
	rootCmd.Flags().SortFlags = true

	return rootCmd
}

// Execute ...
func Execute() {
	if err := RootCmd(&RootConfig{
		Magick: internal.NewMagickClient(),
	}).Execute(); err != nil {
		logger.Println("Error when trying to execute", err)
		os.Exit(1)
	}
}
