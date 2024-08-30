package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	makeconfig "github.com/sitnikovik/go-grpc-api-template/internal/cli/make/config"
)

var rootCmd = &cobra.Command{
	Use:   "warp",
	Short: "Warp is a CLI tool to build your gRPC API service quickly",
	RunE:  root,
}

func init() {
	rootCmd.AddCommand(makeconfig.Cmd)
}

func Execute() {
	rootCmd.Execute()
}

func root(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello, Warp!")
	return nil
}
