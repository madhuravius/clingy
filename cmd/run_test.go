package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"clingy/internal"
)

func TestRunCmdExecuteSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	mockTools.MagickClientImpl.EXPECT().CaptureWindow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	files, err := os.ReadDir("./test_data")
	if err != nil {
		t.Fatalf("Error in reading test_data directory's files %s", err.Error())
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "will_pass") {
			output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"run", "-i", fmt.Sprintf("./test_data/%s", file.Name())})
			assert.Contains(t, output, "Completed clingy run, generating report.")
		}
	}
}

func TestRunCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, Magick: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"run", "--help"})
	assert.Contains(t, output, "Run clingy")
}
