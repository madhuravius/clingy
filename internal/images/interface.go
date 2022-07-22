//go:generate mockgen -source=$GOFILE -destination=../mock/mock_images.go -package=mock
package images

import "log"

// ImageProcessingImpl - used to interact with imagemagick client
type ImageProcessingImpl interface {
	CaptureWindow(logger *log.Logger, buildDirectory string, screenshotName string, screenshotExtension string) (string, error)
	AddLabelToImage(logger *log.Logger, label string, imagePath string) error
	AddDescriptionToImage(logger *log.Logger, description string, imagePath string) error
}
