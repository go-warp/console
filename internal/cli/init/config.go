package init

import (
	"github.com/spf13/cobra"

	"github.com/sitnikovik/go-grpc-api-template/internal/cli/helper/input"
)

type config struct {
	destination string
}

func parseConfig(cmd *cobra.Command) config {
	return config{
		destination: input.GetCmdFlag(cmd, "dest"),
	}
}
