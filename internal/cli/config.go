package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/input"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output/colorize"
)

// configVar is a struct that represents a config variable
type configVar struct {
	Name string // Name of the config variable
	Type string // Type of the config variable
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the config of Warp",
	RunE: func(cmd *cobra.Command, args []string) error {

		vars := []configVar{}
		defer func() {
			if len(vars) > 0 {
				fmt.Printf("\n%s Config has been done with variables:\n", colorize.Green("OK!"))
				for _, v := range vars {
					fmt.Printf(
						"  - Name: %s, Type: %s\n",
						colorize.Yellow(v.Name), colorize.Cyan(v.Type),
					)
				}
			} else {
				output.PrintError("No variables have been added")
			}
		}()

		for {
			v, err := readConfigVar()

			if err != nil {
				if err.Error() == "empty name" {
					break
				}
				output.PrintError(err.Error())
				continue
			}

			fmt.Printf("OK! Name: %s, type: %s\n\n", v.Name, v.Type)
			vars = append(vars, v)
		}

		return nil
	},
}

func readConfigVar() (configVar, error) {
	question := "Specify the variable name: "
	name := input.ReadString(question)
	if name == "" {
		return configVar{}, errors.New("empty name")
	}

	question = fmt.Sprintf("Specify the env type %s:", colorize.Red("[string, int, bool]"))
	envType := input.ReadString(question)
	if !isEnvTypeValid(envType) {
		return configVar{}, errors.New("invalid env type")
	}

	return configVar{Name: name, Type: envType}, nil
}

func isEnvTypeValid(envType string) bool {
	switch envType {
	case "string", "int", "bool":
		return true
	default:
		return false
	}
}
