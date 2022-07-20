//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock
package internal

import "os"

// ExitToolsImpl - used to interact with exit mechanics in a client
type ExitToolsImpl interface {
	Exit(code int)
}

// magickClient - simple struct mainly for testing purposes
type exitTools struct{}

// NewExitToolsClient - generates an interface for reuse
func NewExitToolsClient() ExitToolsImpl {
	return exitTools{}
}

// Exit - wrapper for os exits and/or to do additional logging/telemetry/cleanup during exits
func (e exitTools) Exit(code int) {
	os.Exit(code)
}
