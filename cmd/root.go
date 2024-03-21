package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/list"
	"github.com/w-haibara/kotan/cmd/start"
	"github.com/w-haibara/kotan/cmd/stop"
)

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "kotan",
		Short: "---",
		Long:  "--- --- --- ---",
	}

	cmd.AddCommand(list.Cmd())
	cmd.AddCommand(start.Cmd())
	cmd.AddCommand(stop.Cmd())

	return &cmd
}
