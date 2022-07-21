package internal

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"

	"clingy/internal/mock"
)

type TestInterfaceHousing struct {
	Ctrl             *gomock.Controller
	ExitClientsImpl  *mock.MockExitToolsImpl
	MagickClientImpl *mock.MockImageProcessingImpl
}

// GenerateMockInterfacesForClingy - generates an easy to use struct with helper methods for testing
func GenerateMockInterfacesForClingy(t *testing.T) *TestInterfaceHousing {
	ctrl := gomock.NewController(t)

	return &TestInterfaceHousing{
		Ctrl:             ctrl,
		ExitClientsImpl:  mock.NewMockExitToolsImpl(ctrl),
		MagickClientImpl: mock.NewMockImageProcessingImpl(ctrl),
	}
}

// ExecCobraCmdAndReturnString - executes a cobra command, stores it in buffer, and returns it back out
func ExecCobraCmdAndReturnString(t *testing.T, cmd *cobra.Command, args []string) string {
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs(args)
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}
