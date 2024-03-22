package start

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/client"
	"github.com/w-haibara/kotan/daemon"
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
	resp, err := client.StartUnit(daemon.StartUnitReq{
		Name: args[0],
	})
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
