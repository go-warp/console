package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/input"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output/colorize"
	stringsHelper "github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/strings"
)

const (
	// pathToEnvFile is a path to the .env file
	pathToEnvFile = ".env"
	// pathToConfigFile is a path to the config.go file
	pathToConfigFile = "config.go"

	// configVarTypeString is a string type of the config variable
	configVarTypeString = "string"
	// configVarTypeInt is an int type of the config variable
	configVarTypeInt = "int"
	// configVarTypeBool is a bool type of the config variable explained
	configVarTypeBool = "bool"
)

var (
	errConfigVarEmptyName   = errors.New("empty config var name")
	errConfigVarInvalidType = errors.New("invalid config var type")
)

// configVar is a struct that represents a config variable
type configVar struct {
	Name    string      // Name of the config variable thar will be used in the code
	EnvName string      // Name of the environment variable that will be stored in the .env file
	Usage   string      // Usage of the config variable. Will be used in getter func comment
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

			makeConfig(vars)
		}()

		for {
			v, err := readConfigVar()
			if err != nil {
				if errors.Is(err, errConfigVarEmptyName) { // TODO: check with error variable
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
		return configVar{}, errConfigVarEmptyName
	}

	question = fmt.Sprintf(
		"Specify the env type [%s, %s, %s]:",
		configVarTypeString, configVarTypeInt, configVarTypeBool,
	)
	envType := input.ReadString(question)
	if !isEnvTypeValid(envType) {
		return configVar{}, errConfigVarInvalidType
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

func makeConfig(vars []configVar) error {
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

	// Add the variables
	for _, v := range vars {
		s := fmt.Sprintf("%s=%v\n", strings.ToUpper(v.Name), v.Value)
		sb.WriteString(s)
	}

	// Create the file
	err := output.MakeFile(pathToEnvFile, []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"Created file %s\n",
		colorize.Cyan(pathToEnvFile),
	)
	return nil
}

// makeGoConfigFile creates the config.go file with the specified variables
func makeGoConfigFile(vars []configVar) error {
	sb := strings.Builder{}

	// Set the disclaimer
	sb.WriteString(disclaimer())

	// Package
	sb.WriteString("package config\n\n")

	// Imports
	sb.WriteString("import (\n")
	sb.WriteString("\t\"os\"\n")
	sb.WriteString("\t\"strconv\"\n")
	sb.WriteString(")\n\n")

	// Config struct
	sb.WriteString("// Config is a struct that represents the configuration\n")
	sb.WriteString("type Config struct {\n")
	for _, v := range vars {
		s := fmt.Sprintf(
			"\t%s %s\n",
			stringsHelper.ToCamelCase(v.Name), v.Type,
		)
		sb.WriteString(s)
	}
	sb.WriteString("}\n\n")

	// Constructor
	sb.WriteString("// NewConfig creates a new Config instance\n")
	sb.WriteString("func NewConfig() *Config {\n")
	sb.WriteString("\tc := &Config{}\n")
	sb.WriteString("\n")
	for _, v := range vars {
		switch v.Type {
		case configVarTypeString:
			s := fmt.Sprintf(
				"\tc.%s = os.Getenv(\"%s\")\n",
				stringsHelper.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case configVarTypeInt:
			s := fmt.Sprintf(
				"\tc.%s, _ = strconv.Atoi(os.Getenv(\"%s\"))\n",
				stringsHelper.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case configVarTypeBool:
			s := fmt.Sprintf(
				"\tc.%s, _ = strconv.ParseBool(os.Getenv(\"%s\"))\n",
				stringsHelper.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		}
	}
	sb.WriteString("\n")
	sb.WriteString("\treturn c\n")
	sb.WriteString("}\n")

	// Getters
	for _, v := range vars {
		s := fmt.Sprintf(
			"\n// %s returns the %s value\n",
			stringsHelper.ToPascalCase(v.Name), v.Name,
		)
		sb.WriteString(s)

		s = fmt.Sprintf(
			"func (c *Config) %s() %s {\n\treturn c.%s\n}\n",
			stringsHelper.ToPascalCase(v.Name), v.Type, stringsHelper.ToCamelCase(v.Name),
		)
		sb.WriteString(s)
	}

	// Create the file
	err := output.MakeFile(pathToConfigFile, []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"Created file %s\n",
		colorize.Cyan(pathToConfigFile),
	)

	if err := fixGoImports(pathToConfigFile); err != nil {
		output.PrintError("failed to fix go imports")
	}

	return nil
}
