package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// newCleanCmd - clean temporary paths
func (r *RootConfig) newCleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "clean",
		Short:  "Clean clingy",
		PreRun: initRunWithoutArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Cleaning clingy generated files")

			subPaths, err := ioutil.ReadDir(outputPath)
			if err != nil {
				logger.Println("Unable to read build directory for cleaning", err)
				r.ExitTools.Exit(1)
			}

			for _, subPath := range subPaths {
				if subPath.Name() == ".gitkeep" {
					continue
				}
				err := os.RemoveAll(path.Join([]string{outputPath, subPath.Name()}...))
				if err != nil {
					logger.Println("Unable to clean up normal build path", err)
					r.ExitTools.Exit(1)
				}
			}
			cmd.Println("Finished cleaning build paths.")
		},
	}
}
