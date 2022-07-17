package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd - simple version Command to print
var versionCmd = &cobra.Command{
	Use:    "version",
	Short:  "Print the version number of clingy",
	PreRun: initRunWithoutArtifactDirectoryCreate,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
