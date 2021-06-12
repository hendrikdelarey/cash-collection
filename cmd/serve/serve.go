package serve

import (
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/hendrikdelarey/cash-collection/pkg/app"
)

// Command root command for starting/serving the RESTful part of the service
func Command(ctx context.Context, logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "serve [args]",
		Short: "Start the API up and prepare to receive HTTP requests",
		Run: func(cmd *cobra.Command, args []string) {
			app.Run(ctx, logger)
		},
	}
}