package lib

import (
	"log"
	"strings"
	"testing"
)

func TestParseClingyFile(t *testing.T) {
	type args struct {
		logger   *log.Logger
		fileName string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrPartial string
	}{
		{
			name: "test expected failure on 00_will_fail.yaml",
			args: args{
				logger:   log.Default(),
				fileName: "../cmd/test_data/00_will_fail.yaml",
			},
			wantErr:        true,
			wantErrPartial: "unable to process template, no steps",
		},
		{
			name: "test pass case on 01_basic_flow.yaml",
			args: args{
				logger:   log.Default(),
				fileName: "../cmd/test_data/01_basic_flow.yaml",
			},
			wantErr:        false,
			wantErrPartial: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseClingyFile(tt.args.logger, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseClingyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr && !strings.Contains(err.Error(), tt.wantErrPartial) {
				t.Fatalf("ParseClingyFile() error = %v, partial missing %v", err.Error(), tt.wantErrPartial)
			}
		})
	}
}
