package config

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output"
)

const (
	// pathToEnvFile is a path to the .env file
	pathToEnvFile = ".env"
	// pathToConfigFile is a path to the config.go file
	pathToConfigFile = "config.go"

	// variableTypeString is a string type of the config variable
	variableTypeString = "string"
	// variableTypeInt is an int type of the config variable
	variableTypeInt = "int"
	// variableTypeBool is a bool type of the config variable explained
	variableTypeBool = "bool"
)

var (
	errvariableEmptyName   = errors.New("empty config var name")
	errvariableInvalidType = errors.New("invalid config var type")
)

// NewCommand creates a command that is used to configure the Warp
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "make:config",
		Short: "Make an env config file and its Go representation for the project.",
		RunE:  run,
	}
}

func run(cmd *cobra.Command, args []string) error {
	vars := []variable{}
	defer func() {
		if len(vars) == 0 {
			output.PrintError("No variables have been added")
			return
		}

		makeConfig(vars)
	}()

	for {
		v, err := readVariable()
		if err != nil {
			if errors.Is(err, errvariableEmptyName) {
				break
			}
			output.PrintError(err.Error())
			continue
		}

		fmt.Printf("OK! Name: %s, type: %s\n\n", v.Name, v.Type)
		vars = append(vars, v)
	}

	return nil
}
