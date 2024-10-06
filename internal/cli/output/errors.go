package output

import (
	"fmt"

	"github.com/go-warp/console/internal/cli/output/colorize"
)

// PrintError prints an error message to the console
func PrintError(s string) {
	fmt.Printf("\n%s: %s\n\n", colorize.Red("ERROR"), s)
}
