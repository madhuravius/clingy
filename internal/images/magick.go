package images

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// magickClient - simple struct mainly that scopes methods to imagemagick for linux users
type magickClient struct{}

// NewMagickClient - generates an interface for reuse
func NewMagickClient() ImageProcessingImpl {
	return magickClient{}
}

// CaptureWindow - executes command and captures window contents
func (m magickClient) CaptureWindow(
	logger *log.Logger,
	buildDirectory string,
	screenshotName string,
	screenshotExtension string,
) (string, error) {
	// typical execution path - magick import -window $WINDOWID ./images/{screenshot}
	logger.Println("Taking screenshot", fmt.Sprintf("WINDOWID - %s", os.Getenv("WINDOWID")))
	expectedPath := fmt.Sprintf("%s/%s%s", buildDirectory, screenshotName, screenshotExtension)
	logger.Println("Saving to path", expectedPath)

	imageCommand := exec.Command(
		"magick",
		"import",
		"-window",
		os.Getenv("WINDOWID"),
		expectedPath,
	)
	if err := imageCommand.Run(); err != nil {
		return "", err
	}

	logger.Println("Saved to path", expectedPath)
	return expectedPath, nil
}

// AddLabelToImage - add title text to image
func (m magickClient) AddLabelToImage(logger *log.Logger, label string, imagePath string) error {
	return addLabelToImage(logger, label, imagePath)
}

// AddDescriptionToImage - add description text to image
func (m magickClient) AddDescriptionToImage(logger *log.Logger, description string, imagePath string) error {
	return addDescriptionToImage(logger, description, imagePath)
}
