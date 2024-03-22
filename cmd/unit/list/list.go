package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/client"
	"github.com/w-haibara/kotan/daemon"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List unit files",
		Long:  "--- --- --- ---",
		RunE:  list,
	}

	return &cmd
}

func list(cmd *cobra.Command, args []string) error {
	resp, err := client.ListUnit(daemon.ListUnitReq{})
	if err != nil {
		return err
	}

	// TODO: Pretty print units
	for _, u := range resp.UnitInfo {
		fmt.Println(u.Name)
	}

	return nil
}
