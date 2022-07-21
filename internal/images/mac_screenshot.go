package images

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// macScreenshotClient - simple struct that scopes interface methods for Mac clients
type macScreenshotClient struct{}

// NewMacScreenshotClient - generates an interface for reuse
func NewMacScreenshotClient() ImageProcessingImpl {
	return macScreenshotClient{}
}

// CaptureWindow - executes command and captures window contents
func (m macScreenshotClient) CaptureWindow(
	logger *log.Logger,
	buildDirectory string,
	screenshotName string,
	screenshotExtension string,
) (string, error) {
	// typical execution path - magick import -window $WINDOWID ./images/{screenshot}
	logger.Println("Taking screenshot", fmt.Sprintf("WINDOW_NAME - %s", os.Getenv("WINDOW_NAME")))
	expectedPath := fmt.Sprintf("%s/%s%s", buildDirectory, screenshotName, screenshotExtension)
	logger.Println("Saving to path", expectedPath)
	// screenshot -w on_screen_only WINDOW_NAME
	imageCommand := exec.Command(
		"screenshot",
		"-w",
		"on_screen_only",
		os.Getenv("WINDOW_NAME"),
		"-f",
		expectedPath,
	)
	if err := imageCommand.Run(); err != nil {
		return "", err
	}

	logger.Println("Saved to path", expectedPath)
	return expectedPath, nil
}

// AddLabelToImage - add title text to image
func (m macScreenshotClient) AddLabelToImage(logger *log.Logger, label string, imagePath string) error {
	return addLabelToImage(logger, label, imagePath)
}

// AddDescriptionToImage - add description text to image
func (m macScreenshotClient) AddDescriptionToImage(logger *log.Logger, description string, imagePath string) error {
	return addDescriptionToImage(logger, description, imagePath)
}
