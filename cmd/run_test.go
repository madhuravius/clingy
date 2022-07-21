package cmd

import (
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestRunCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	mockTools.MagickClientImpl.EXPECT().CaptureWindow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(4)
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"run", "-i", "./test_data/01_basic_flow.yaml"})
	assert.Contains(t, output, "Completed clingy run, generating report.")
}

func TestRunCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"run", "--help"})
	assert.Contains(t, output, "Run clingy")
}
