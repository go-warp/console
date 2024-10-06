package config

import (
	"fmt"

	"github.com/go-warp/console/internal/cli/input"
)

const (
	// variableTypeString is a string type of the config variable
	variableTypeString = "string"
	// variableTypeInt is an int type of the config variable
	variableTypeInt = "int"
	// variableTypeBool is a bool type of the config variable explained
	variableTypeBool = "bool"
)

// variable is a struct that represents a config variable
type variable struct {
	Name    string      // Name of the config variable thar will be used in the code
	EnvName string      // Name of the environment variable that will be stored in the .env file
	Usage   string      // Usage of the config variable. Will be used in getter func comment
	Type    string      // Type of the config variable
	Value   interface{} // Current value of the config variable
}

// readVariable reads the config variable from the user
func readVariable() (variable, error) {
	question := "Specify the variable name: "
	name := input.ReadString(question)
	if name == "" {
		return variable{}, errvariableEmptyName
	}

	question = fmt.Sprintf(
		"Specify the env type [%s, %s, %s]:",
		variableTypeString, variableTypeInt, variableTypeBool,
	)
	envType := input.ReadString(question)
	if envType == "" {
		envType = variableTypeString
	}
	if !isEnvTypeValid(envType) {
		return variable{}, errvariableInvalidType
	}

	question = "If you need, specify the inital value:"
	value := input.ReadString(question)

	return variable{
		Name:  name,
		Type:  envType,
		Value: value,
	}, nil
}

// isEnvTypeValid checks if the env type is valid
func isEnvTypeValid(envType string) bool {
	switch envType {
	case variableTypeString,
		variableTypeInt,
		variableTypeBool:
		return true
	default:
		return false
	}
}
