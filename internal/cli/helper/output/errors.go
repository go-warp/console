package output

import (
	"fmt"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/output/colorize"
)

// PrintError prints an error message to the console
func PrintError(s string) {
	fmt.Printf("\n%s: %s\n\n", colorize.Red("ERROR"), s)
}
