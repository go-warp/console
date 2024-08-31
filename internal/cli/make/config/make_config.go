package config

import (
	"fmt"
	"strings"

	gofile "github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/gofile"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output"
	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output/colorize"
	stringsHelper "github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/strings"
)

const (
	// pathToEnvFile is a path to the .env file
	pathToEnvFile = ".env"
	// pathToConfigFile is a path to the config.go file
	pathToConfigFile = "config.go"
)

// makeConfig creates config files with the specified variables
func makeConfig(vars []variable) error {
	if err := makeEnvFile(vars); err != nil {
		return err
	}
	if err := makeGoConfigFile(vars); err != nil {
		return err
	}

	return nil
}

// makeEnvFile creates the .env file with the specified variables
func makeEnvFile(vars []variable) error {
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
func makeGoConfigFile(vars []variable) error {
	sb := strings.Builder{}

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
		case variableTypeString:
			s := fmt.Sprintf(
				"\tc.%s = os.Getenv(\"%s\")\n",
				stringsHelper.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case variableTypeInt:
			s := fmt.Sprintf(
				"\tc.%s, _ = strconv.Atoi(os.Getenv(\"%s\"))\n",
				stringsHelper.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case variableTypeBool:
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

	if err := gofile.FixGoimports(pathToConfigFile); err != nil {
		output.PrintError("failed to fix go imports")
	}

	return nil
}
