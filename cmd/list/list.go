package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/unit"
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
	units, err := unit.List()
	if err != nil {
		return err
	}

	fmt.Println(units)

	return nil
}
