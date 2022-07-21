package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestInitCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	defer func() {
		if err := os.RemoveAll("../output/.test.clingy.yaml"); err != nil {
			t.Fatal(err)
		}
	}()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"init", "-i", "../output/.test.clingy.yaml"})
	assert.Contains(t, output, "Finished writing template to")
}

func TestInitCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"init", "--help"})
	assert.Contains(t, output, "instantiate a .clingy.yaml (or input name of your choice) to pwd")
}
