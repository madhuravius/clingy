package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newVersionCmd - simple version Command to print
func (r *RootConfig) newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "version",
		Short:  "Print the version number of clingy",
		PreRun: initRunWithoutArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}
