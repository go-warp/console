package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the config of Warp",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Doing some config stuff...\n\n")

		name := helper.ReadString("Specify the env name: ")

		envType := helper.ReadString("Specify the env type [string, int, bool]: ")
		if !isEnvTypeValid(envType) {
			return fmt.Errorf("invalid env type")
		}

		fmt.Printf("OK! Name: %s, type: %s\n", name, envType)

		return nil
	},
}

func isEnvTypeValid(envType string) bool {
	switch envType {
	case "string", "int", "bool":
		return true
	default:
		return false
	}
}
