package daemon

import (
	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/daemon"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "daemon",
		Short: "---",
		Long:  "--- --- --- ---",
		RunE:  run,
	}

	return &cmd
}

func run(cmd *cobra.Command, args []string) error {
	if err := daemon.Run(); err != nil {
		return err
	}

	return nil
}
