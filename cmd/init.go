package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// templateFile - yaml formatted
const templateFile = `label: clingy flow
steps:
- label: start
  description: starting clingy flow
  command: echo 
  args:
    - "Starting"`

// newInitCmd - inits a .clingy.yaml for use in the current path
func (r *RootConfig) newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:    "init",
		Short:  "instantiate a .clingy.yaml for use in the cwd",
		PreRun: initRunWithoutArtifactDirectoryCreate,
		Run: func(cmd *cobra.Command, args []string) {
			// check current path to determine if needing to write file
			logger.Println("Checking if inputFile exists already", inputFile)
			fileInfo, err := os.Stat(inputFile)
			if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
				logger.Println("Error in os stat for file info", err)
				fmt.Println("Error in looking up path to write the file")
				os.Exit(1)
			}
			if fileInfo != nil {
				logger.Println("File found: ", fileInfo.Name())
				fmt.Println("File already exists!")
				os.Exit(1)
			}

			// if it doesn't exist, go ahead and write the default template
			if err := os.WriteFile(inputFile, []byte(templateFile), 0644); err != nil {
				logger.Println("Error in writing file", err)
				fmt.Println("Error in trying to write template file")
				os.Exit(1)
			}
		},
	}

}
