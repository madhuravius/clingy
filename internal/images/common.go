package images

import (
	"log"
	"os/exec"
)

// addLabelToImage - adds a label to an image with imagemagick
func addLabelToImage(logger *log.Logger, imagePath, label string) error {
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

// addDescriptionToImage
func addDescriptionToImage(logger *log.Logger, imagePath, description string) error {
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
