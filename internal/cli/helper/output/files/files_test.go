package files

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testdataDir = "testdata/"

func TestCreate(t *testing.T) {
	// Test creates files, so we can't run it in parallel
	type args struct {
		path    string
		content []byte
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
				path:    testdataDir + "create.txt",
				content: []byte("Hello, World!\n"),
			},
			fileToCompare: testdataDir + "create_expected.txt",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Create(tt.args.path, tt.args.content)
			defer os.Remove(tt.args.path) // Clean up

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			// Compare the files
			bbGot, err := os.ReadFile(tt.args.path)
			require.NoError(t, err)
			bbWant, err := os.ReadFile(tt.fileToCompare)
			require.NoError(t, err)
			require.Equal(t, bbWant, bbGot)
		})
	}
}

func TestFixGoimports(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name          string
		args          args
		fileToCompare string
		fileToCopy    string
		wantErr       bool
	}{
		{
			name: "ok",
			args: args{
				path: testdataDir + "fix_goimports.go",
			},
			fileToCompare: testdataDir + "fix_goimports_expected.txt",
			fileToCopy:    testdataDir + "fix_goimports.txt",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the file to work with
			bb, err := os.ReadFile(tt.fileToCompare)
			require.NoError(t, err)
			err = os.WriteFile(tt.args.path, bb, 0644)
			require.NoError(t, err)
			defer os.Remove(tt.args.path) // Clean up

			// Call the function we want to test
			err = FixGoimports(tt.args.path)

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			// Compare the files
			bbWant, err := os.ReadFile(tt.fileToCompare)
			require.NoError(t, err)
			bbGot, err := os.ReadFile(tt.args.path)
			require.NoError(t, err)

			require.Equal(t, bbWant, bbGot)
		})
	}
}
