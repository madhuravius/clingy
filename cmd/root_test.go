package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"clingy/internal"
)

func TestRootCmdExecute(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd)
	assert.Contains(t, output, "clingy is a tool to test and capture CLI flows")
}
