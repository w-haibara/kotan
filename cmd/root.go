package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/w-haibara/kotan/cmd/daemon"
	"github.com/w-haibara/kotan/cmd/unit"
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

	cmd.AddCommand(daemon.Cmd())
	cmd.AddCommand(unit.Cmd())

	return &cmd
}
