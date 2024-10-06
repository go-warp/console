package init

import (
	"github.com/spf13/cobra"

	"github.com/go-warp/console/internal/cli/input"
)

type config struct {
	destination string
}

func parseConfig(cmd *cobra.Command) config {
	return config{
		destination: input.GetCmdFlag(cmd, "dest"),
	}
}
