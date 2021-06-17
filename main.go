package main

import (
	"context"
	"fmt"
	"github.com/hendrikdelarey/cash-collection/pkg/app"

	"go.uber.org/zap"

	"github.com/hendrikdelarey/cash-collection/cmd/serve"
	"github.com/spf13/cobra"
)

const serviceName string = "cash-collection"

func main() {
	ctx := context.Background()
	cmd := &cobra.Command{Use: serviceName}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprint("failed to load logger"))
	}

	cmd.AddCommand(serve.Command(ctx, logger))
	if err := cmd.Execute(); err != nil {
		logger.Fatal("fatal error", zap.Error(err))
	}

	app.Run(ctx, logger)
}
