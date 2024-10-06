package config

import (
	"fmt"
	"strings"

	"github.com/sitnikovik/stringo/cases"

	"github.com/go-warp/console/internal/cli/output"
	"github.com/go-warp/console/internal/cli/output/colorize"
	"github.com/go-warp/console/internal/cli/output/files"
)

const (
	// pathToEnvFile is a path to the .env file
	pathToEnvFile = ".env"
	// pathToConfigFile is a path to the config.go file
	pathToConfigFile = "internal/config/config.go"
)

// makeConfig creates config files with the specified variables
func makeConfig(vars []variable) error {
	// TODO: think how to test this method properly
	if err := makeEnvFile(pathToEnvFile, vars); err != nil {
		return err
	}
	if err := makeGoConfigFile(pathToConfigFile, vars); err != nil {
		return err
	}

	return nil
}

// makeEnvFile creates the .env file with the specified variables
func makeEnvFile(path string, vars []variable) error {
	sb := strings.Builder{}

	// Add the variables
	for _, v := range vars {
		s := fmt.Sprintf("%s=%v\n", strings.ToUpper(v.Name), v.Value)
		sb.WriteString(s)
	}

	// Create the file
	err := files.Create(path, []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"Created file %s\n",
		colorize.Cyan(path),
	)
	return nil
}

// makeGoConfigFile creates the config.go file with the specified variables
func makeGoConfigFile(path string, vars []variable) error {
	sb := strings.Builder{}

	// Package
	sb.WriteString("package config\n\n")

	// Imports
	sb.WriteString("import (\n")
	sb.WriteString("\t\"os\"\n")
	sb.WriteString("\t\"strconv\"\n")
	sb.WriteString("\n")
	sb.WriteString("\t\"github.com/joho/godotenv\"\n")
	sb.WriteString(")\n\n")

	// Init
	sb.WriteString("\n")
	sb.WriteString("// Init initializes the configuration\n")
	sb.WriteString("func Init(_ context.Context) error {\n")
	sb.WriteString("\treturn godotenv.Load(\".env\")\n")
	sb.WriteString("}\n")

	// Config struct
	sb.WriteString("// Config is a struct that represents the configuration\n")
	sb.WriteString("type Config struct {\n")
	for _, v := range vars {
		s := fmt.Sprintf(
			"\t%s %s\n",
			cases.ToCamelCase(v.Name), v.Type,
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
				cases.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case variableTypeInt:
			s := fmt.Sprintf(
				"\tc.%s, _ = strconv.Atoi(os.Getenv(\"%s\"))\n",
				cases.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		case variableTypeBool:
			s := fmt.Sprintf(
				"\tc.%s, _ = strconv.ParseBool(os.Getenv(\"%s\"))\n",
				cases.ToCamelCase(v.Name), v.Name,
			)
			sb.WriteString(s)
		}
	}
	sb.WriteString("\n")
	sb.WriteString("\treturn c\n")
	sb.WriteString("}\n")

	// Getters
	for _, v := range vars {
		pascalCase := cases.ToPascalCase(v.Name)
		s := fmt.Sprintf(
			"\n// %s returns the %s value\n",
			pascalCase, v.Name,
		)
		sb.WriteString(s)

		s = fmt.Sprintf(
			"func (c *Config) %s() %s {\n\treturn c.%s\n}\n",
			pascalCase,
			v.Type,
			cases.ToCamelCase(v.Name),
		)
		sb.WriteString(s)
	}

	// Create the file
	err := files.Create(path, []byte(sb.String()))
	if err != nil {
		return err
	}

	fmt.Printf(
		"Created file %s\n",
		colorize.Cyan(path),
	)

	if err := files.FixGoimports(path); err != nil {
		output.PrintError("failed to fix go imports")
	}

	return nil
}
