package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestRootCmdExecuteInstantiation(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{})
	assert.Contains(t, output, "clingy is a tool to test and capture CLI flows")
}

func TestRootCmdExecuteHelp(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"--help"})
	assert.Contains(t, output, "clingy is a tool to test and capture CLI flows")
}
