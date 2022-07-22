package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestRootCmdExecuteInstantiation(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{})
	assert.Contains(t, output, "clingy is a tool to test and capture CLI flows")
}

func TestRootCmdExecuteHelp(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"--help"})
	assert.Contains(t, output, "clingy is a tool to test and capture CLI flows")
}

func TestRootCmdExecuteInvalidCommand(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"will-exit-1-does-not-exist"})
	cmdErr := cmd.Execute()
	out, _ := ioutil.ReadAll(b)
	assert.NotNil(t, cmdErr)
	assert.Contains(t, string(out), "Error: unknown command")
}
