package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_readVariable(t *testing.T) {
	// Test uses Stdin, so we can't run it in parallel
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "ok",
			input:   "FOO\nstring\n",
			wantErr: false,
		},
		{
			name:    "err on empty name",
			input:   "\n",
			wantErr: true,
		},
		{
			name:    "err on invalid type",
			input:   "FOO\n1\n",
			wantErr: true,
		},
		{
			name:    "ok on empty type return string type variable",
			input:   "FOO\n\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save the original os.Stdin
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()

			// Create a pipe to simulate stdin
			r, w, err := os.Pipe()
			require.NoError(t, err)

			// Write the test input to the pipe
			_, err = w.Write([]byte(tt.input))
			require.NoError(t, err)
			w.Close()

			// Replace os.Stdin with the read end of the pipe
			os.Stdin = r

			// Call the function you want to test
			v, err := readVariable()

			// Check the results
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, v)
			}
		})
	}
}

func Test_isEnvTypeValid(t *testing.T) {
	t.Parallel()
	type args struct {
		envType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok string",
			args: args{envType: variableTypeString},
			want: true,
		},
		{
			name: "ok int",
			args: args{envType: variableTypeInt},
			want: true,
		},
		{
			name: "ok bool",
			args: args{envType: variableTypeBool},
			want: true,
		},
		{
			name: "err integer",
			args: args{envType: "foo"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isEnvTypeValid(tt.args.envType)

			require.Equal(t, tt.want, got)
		})
	}
}
