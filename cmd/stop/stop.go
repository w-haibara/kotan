package stop

import (
	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/unit"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "stop",
		Short: "Stop unit",
		Long:  "--- --- --- ---",
		Args:  cobra.ExactArgs(1),
		RunE:  stop,
	}

	return &cmd
}

func stop(cmd *cobra.Command, args []string) error {
	unit, err := unit.Load(args[0])
	if err != nil {
		return err
	}

	if err := unit.Stop(); err != nil {
		return err
	}

	return nil
}
