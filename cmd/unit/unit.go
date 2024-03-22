package unit

import (
	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/unit/list"
	"github.com/w-haibara/kotan/cmd/unit/start"
	"github.com/w-haibara/kotan/cmd/unit/stop"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "unit",
		Short: "---",
		Long:  "--- --- --- ---",
	}

	cmd.AddCommand(list.Cmd())
	cmd.AddCommand(start.Cmd())
	cmd.AddCommand(stop.Cmd())

	return &cmd
}
