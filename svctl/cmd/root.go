package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewCLI() *cobra.Command {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	cmd := &cobra.Command{
		Use: "svctl",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(newServeCmd(logger))
	return cmd
}
