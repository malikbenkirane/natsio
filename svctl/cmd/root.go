package cmd

import "github.com/spf13/cobra"

func NewCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use: "svctl",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return cmd
}
