//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// MagickClientImpl - used to interact with imagemagick client
type MagickClientImpl interface {
	CaptureWindow(logger *log.Logger, buildDirectory string, screenshotName string, screenshotExtension string) (string, error)
	AddLabelToImage(logger *log.Logger, label string, imagePath string) error
	AddDescriptionToImage(logger *log.Logger, description string, imagePath string) error
}

// magickClient - simple struct mainly for testing purposes
type magickClient struct{}

// NewMagickClient - generates an interface for reuse
func NewMagickClient() MagickClientImpl {
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
		return err
	}
	logger.Println("Saved label to image path", imagePath)

	return nil
}

// AddDescriptionToImage - add description text to image
func (m magickClient) AddDescriptionToImage(logger *log.Logger, description string, imagePath string) error {
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
		return err
	}
	logger.Println("Saved description to image path", imagePath)

	return nil
}
