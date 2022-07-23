package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
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

func TestRunCmdExecuteFailures(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	mockTools.MagickClientImpl.EXPECT().CaptureWindow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	files, err := os.ReadDir("./test_data")
	if err != nil {
		t.Fatalf("Error in reading test_data directory's files %s", err.Error())
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "will_fail") {
			mockTools.ExitClientsImpl.EXPECT().Exit(1) // this is important, make sure we expect a failure
			b := new(bytes.Buffer)
			cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs([]string{"run", "-o", "../output", "-i", fmt.Sprintf("./test_data/%s", file.Name())})
			_ = cmd.Execute()
			out, _ := ioutil.ReadAll(b)
			assert.Contains(t, string(out), "Error ")
		}
	}
}

func TestRunCmdHelpSuccess(t *testing.T) {
	mockTools := internal.GenerateMockInterfacesForClingy(t)
	defer mockTools.Ctrl.Finish()

	cmd := RootCmd(&RootConfig{ExitTools: mockTools.ExitClientsImpl, ImageTools: mockTools.MagickClientImpl})
	output := internal.ExecCobraCmdAndReturnString(t, cmd, []string{"run", "--help"})
	assert.Contains(t, output, "Run clingy")
}
