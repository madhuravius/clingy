package cmd

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestValidateCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	mockTools.ExitClientsImpl.EXPECT().Exit(gomock.Any())

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"validate", "-i", "./test_data/01_basic_flow.yaml"})
	assert.Contains(t, output, "Completed validation, looks good!")
}

func TestValidateCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"validate", "--help"})
	assert.Contains(t, output, "Validate a clingy.yml file\n\nUsage:")
}
