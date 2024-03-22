package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/unit"
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
	units := unit.List()

	// TODO: Pretty print units
	for _, u := range units {
		fmt.Println(u.Name)
	}

	return nil
}
