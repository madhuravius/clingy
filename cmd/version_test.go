package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestVersionCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"version"})
	assert.Contains(t, output, version)
}

func TestVersionCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"version", "--help"})
	assert.Contains(t, output, "Print the version number")
}
