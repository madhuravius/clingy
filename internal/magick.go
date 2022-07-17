package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// CaptureWindow - captures the terminal window
func CaptureWindow(
	logger *log.Logger,
	buildDirectory string,
	command string,
	commandArgs []string,
	screenshotName string,
) string {
	// typical execution path - magick import -window $WINDOWID ./images/{screenshot}
	logger.Println("Taking screenshot", fmt.Sprintf("WINDOWID - %s", os.Getenv("WINDOWID")))
	expectedPath := fmt.Sprintf("%s/%s.jpg", buildDirectory, screenshotName)
	logger.Println("Saving to path", expectedPath)

	fmt.Println("> ", command, strings.Join(commandArgs, " "))
	output, _ := exec.Command(command, commandArgs...).CombinedOutput() // always allow the command to possibly error
	fmt.Println(string(output))
	logger.Println("Finished executing command", command)
	time.Sleep(1000 * time.Millisecond) // waiting because the parent terminal process may not have finished rendering

	imageCommand := exec.Command(
		"magick",
		"import",
		"-window",
		os.Getenv("WINDOWID"),
		expectedPath,
	)
	if err := imageCommand.Run(); err != nil {
		log.Println("Error in capturing screenshot", err)
		os.Exit(1)
	}

	logger.Println("Saved to path", expectedPath)
	return expectedPath
}

// AddLabelToImage - add title text to image
func AddLabelToImage(logger *log.Logger, label string, imagePath string) {
	// magick 0.jpg  -font "FreeMono" -gravity South -pointsize 30 -fill "yellow" -annotate +0+100 'Caption' 0.jpg
	imageCommand := exec.Command(
		"magick",
		imagePath,
		"-font",
		"FreeMono",
		"-gravity",
		"South",
		"-pointsize",
		"30",
		"-fill",
		"yellow",
		"-annotate",
		"+0+100",
		label,
		imagePath,
	)
	output, err := imageCommand.CombinedOutput()
	logger.Println("Combined output of label insertion", string(output))
	if err != nil {
		os.Exit(1)
	}
	logger.Println("Saved label to image path", imagePath)
}

// AddDescriptionToImage - add description text to image
func AddDescriptionToImage(logger *log.Logger, description string, imagePath string) {
	//  magick 0.jpg  -font "FreeMono" -gravity South -pointsize 16 -fill "yellow" -annotate +0+60 'Description text. ' 0.jpg
	imageCommand := exec.Command(
		"magick",
		imagePath,
		"-font",
		"FreeMono",
		"-gravity",
		"South",
		"-pointsize",
		"16",
		"-fill",
		"yellow",
		"-annotate",
		"+0+60",
		description,
		imagePath,
	)
	output, err := imageCommand.CombinedOutput()
	logger.Println("Combined output of description insertion", string(output))
	if err != nil {
		os.Exit(1)
	}
	logger.Println("Saved description to image path", imagePath)
}
