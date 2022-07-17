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
	version = "v0.1.0"

	// flags
	// debug - enable verbose logging
	debug bool
	// outputPath - a location to dump artifacts/output as a result of a clingy run
	outputPath = "./output"
	// inputFile - a path that contains an input file to digest and run clingy against
	inputFile = "./.clingy.yaml"
)

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
	logger = internal.InitLogger(getOutputPath(), debug)
	if debug {
		internal.InitDestinationDirectory(getOutputPath())
	}
}

// rootCmd - entrypoint for clingy app
var rootCmd = &cobra.Command{
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

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Println("Error when trying to execute", err)
		os.Exit(1)
	}
}

// init ...
func init() {
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logs")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "outputPath", "o", outputPath, "build path that dumps outputs")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "inputFile", "i", inputFile, "inputFile representing a .clingy.yaml")
	rootCmd.Flags().SortFlags = true
}