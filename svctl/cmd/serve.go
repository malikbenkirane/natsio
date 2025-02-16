package cmd

import (
	"fmt"

	"github.com/malikbenkirane/natsio/svctl/internal/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newServeCmd(log *zap.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			svc, err := service.New(log.Named("service"))
			if err != nil {
				return fmt.Errorf("new service: %w", err)
			}
			return svc.Start(cmd.Context())
		},
	}
	return cmd
}
