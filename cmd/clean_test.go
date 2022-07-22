package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestCleanCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"clean", "-o", "../output"})
	assert.Contains(t, output, "Finished cleaning build paths.")
}

func TestCleanCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"clean", "-o", "../output", "--help"})
	assert.Contains(t, output, "Clean clingy")
}
