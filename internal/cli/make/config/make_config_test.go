package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testdataDir = "testdata/"

func Test_makeEnvFile(t *testing.T) {
	// Test creates files, so we can't run it in parallel
	type args struct {
		path string
		vars []variable
	}
	tests := []struct {
		name          string
		args          args
		fileToCompare string
		wantErr       bool
	}{
		{
			name: "ok",
			args: args{
				path: testdataDir + ".env",
				vars: []variable{
					{
						Name:  "foo",
						Type:  variableTypeString,
						Value: "bar",
					},
				},
			},
			fileToCompare: testdataDir + "env_expected.txt",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := makeEnvFile(tt.args.path, tt.args.vars)

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			defer os.Remove(tt.args.path) // Clean up

			require.NoError(t, err)
			require.FileExists(t, tt.args.path)

			// Compare the files
			bbGot, err := os.ReadFile(tt.args.path)
			require.NoError(t, err)
			bbWant, err := os.ReadFile(tt.fileToCompare)
			require.NoError(t, err)
			require.Equal(t, bbWant, bbGot)
		})
	}
}

func Test_makeGoConfigFile(t *testing.T) {
	// Test creates files, so we can't run it in parallel
	type args struct {
		path string
		vars []variable
	}
	tests := []struct {
		name          string
		args          args
		fileToCompare string
		wantErr       bool
	}{
		{
			name: "ok to exists dir",
			args: args{
				path: testdataDir + "config.go",
				vars: []variable{
					{
						Name: "FOO_BAR",
						Type: variableTypeString,
					},
				},
			},
			fileToCompare: testdataDir + "config_go_expected.txt",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := makeGoConfigFile(tt.args.path, tt.args.vars)

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			defer os.Remove(tt.args.path) // Clean up

			require.NoError(t, err)
			require.FileExists(t, tt.args.path)

			// Compare the files
			bbGot, err := os.ReadFile(tt.args.path)
			require.NoError(t, err)
			bbWant, err := os.ReadFile(tt.fileToCompare)
			require.NoError(t, err)
			require.Equal(t, bbWant, bbGot)
		})
	}
}
