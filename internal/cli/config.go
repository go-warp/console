package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/input"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output/colorize"
)

const (
	// pathToEnvFile is a path to the .env file
	pathToEnvFile = ".env"
)

// configVar is a struct that represents a config variable
type configVar struct {
	Name    string      // Name of the config variable thar will be used in the code
	EnvName string      // Name of the environment variable that will be stored in the .env file
	Type    string      // Type of the config variable
	Value   interface{} // Current value of the config variable
}

// configCmd is a command that is used to configure the Warp
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the config of Warp",
	RunE: func(cmd *cobra.Command, args []string) error {

		vars := []configVar{}
		defer func() {
			if len(vars) == 0 {
				output.PrintError("No variables have been added")
				return
			}

			processVars(vars)
		}()

		for {
			v, err := readConfigVar()

			if err != nil {
				if err.Error() == "empty name" { // TODO: check with error variable
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

// readConfigVar reads the config variable from the user
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

	question = "If you need, specify the inital value:"
	value := input.ReadString(question)

	return configVar{
		Name:  name,
		Type:  envType,
		Value: value,
	}, nil
}

// isEnvTypeValid checks if the env type is valid
func isEnvTypeValid(envType string) bool {
	switch envType {
	case "string", "int", "bool":
		return true
	default:
		return false
	}
}

func processVars(vars []configVar) error {
	if err := makeEnvFile(vars); err != nil {
		return err
	}
	if err := makeGoConfigFile(vars); err != nil {
		return err
	}

	return nil
}

// makeEnvFile creates the .env file with the specified variables
func makeEnvFile(vars []configVar) error {
	sb := strings.Builder{}
	for _, v := range vars {
		s := fmt.Sprintf("%s=%v\n", strings.ToUpper(v.Name), v.Value)
		sb.WriteString(s)
	}

	err := output.MakeFile(pathToEnvFile, []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"\n%s Config file has been placed by %s\n",
		colorize.Green("OK!"),
		colorize.Cyan(pathToEnvFile),
	)
	return nil
}

// makeGoConfigFile creates the config.go file with the specified variables
func makeGoConfigFile(vars []configVar) error {
	sb := strings.Builder{}

	sb.WriteString("package config\n\n")

	sb.WriteString("import (\n")
	sb.WriteString("\t\"os\"\n")
	sb.WriteString("\t\"strconv\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString("var (\n")
	for _, v := range vars {
		switch v.Type {
		case "string":
			sb.WriteString(fmt.Sprintf("\t%s = os.Getenv(\"%s\")\n", v.Name, strings.ToUpper(v.Name)))
		case "int":
			sb.WriteString(fmt.Sprintf("\t%s, _ = strconv.Atoi(os.Getenv(\"%s\"))\n", v.Name, strings.ToUpper(v.Name)))
		case "bool":
			sb.WriteString(fmt.Sprintf("\t%s, _ = strconv.ParseBool(os.Getenv(\"%s\"))\n", v.Name, strings.ToUpper(v.Name)))
		}
	}
	sb.WriteString(")\n")

	err := output.MakeFile("config.go", []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"\n%s Config file has been placed by %s\n",
		colorize.Green("OK!"),
		colorize.Cyan("config.go"),
	)
	return nil
}
