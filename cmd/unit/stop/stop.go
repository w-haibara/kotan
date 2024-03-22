package stop

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/client"
	"github.com/w-haibara/kotan/daemon"
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
	resp, err := client.StopUnit(daemon.StopUnitReq{
		Name: args[0],
	})
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
