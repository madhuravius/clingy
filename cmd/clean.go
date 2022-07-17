package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// cleanCmd - clean temporary paths
var cleanCmd = &cobra.Command{
	Use:    "clean",
	Short:  "Clean clingy",
	PreRun: initRunWithoutArtifactDirectoryCreate,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("Cleaning clingy generated files")

		subPaths, err := ioutil.ReadDir(getOutputPath())
		if err != nil {
			logger.Println("Unable to read build directory for cleaning", err)
			os.Exit(1)
		}
		for _, subPath := range subPaths {
			if subPath.Name() == ".gitkeep" {
				continue
			}
			err := os.RemoveAll(path.Join([]string{getOutputPath(), subPath.Name()}...))
			if err != nil {
				logger.Println("Unable to clean up normal build path", err)
				os.Exit(1)
			}
		}
	},
}
