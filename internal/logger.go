package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// InitDestinationDirectory - sets up a destination directory for reuse for images/logs
func InitDestinationDirectory(buildPath string) {
	err := os.MkdirAll(buildPath, 0755)
	if err != nil {
		fmt.Println("Error setting up the build", err)
		os.Exit(1)
	}
}

// InitLogger - inits a debug logger for use if needed
func InitLogger(buildPath string, debug bool) *log.Logger {
	var logger *log.Logger
	if debug {
		file, err := os.OpenFile(
			fmt.Sprintf("%s/logs.txt", buildPath),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666,
		)
		if err != nil {
			log.Fatal(err)
		}

		logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logger = log.Default()
		logger.SetOutput(ioutil.Discard)
	}
	return logger
}
