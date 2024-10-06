package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ReadString reads a string from the user
func ReadString(question string) string {
	fmt.Println(question)

	var input string
	fmt.Scanln(&input)

	return input
}

// GetCmdFlag parse command flags and returns value specified with name
func GetCmdFlag(cmd *cobra.Command, name string) string {
	f := cmd.Flag(name)
	if f == nil {
		return ""
	}

	return f.Value.String()
}
