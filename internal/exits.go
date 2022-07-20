package internal

import "os"

// Exit - wrapper for os exits and/or to do additional logging/telemetry/cleanup during exits
func (e exitTools) Exit(code int) {
	os.Exit(code)
}
