package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd - inits a .clingy.yaml for use in the current path
var initCmd = &cobra.Command{
	Use:    "init",
	Short:  "instantiate a .clingy.yaml for use in the cwd",
	PreRun: initRunWithoutArtifactDirectoryCreate,
	Run: func(cmd *cobra.Command, args []string) {
		// check current path to determine if needing to write file

		// if it doesn't exist, go ahead and write the default template

	},
}
