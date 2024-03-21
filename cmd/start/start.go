package start

import (
	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/unit"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "Start unit",
		Long:  "--- --- --- ---",
		Args:  cobra.ExactArgs(1),
		RunE:  start,
	}

	return &cmd
}

func start(cmd *cobra.Command, args []string) error {
	unit, err := unit.Load(args[0])
	if err != nil {
		return err
	}

	if err := unit.Start(); err != nil {
		return err
	}

	return nil
}
