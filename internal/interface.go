//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package internal

import "log"

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
