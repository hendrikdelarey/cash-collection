package app

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hendrikdelarey/cash-collection/pkg/router"
)

// Run starts the HTTP application server
func Run(ctx context.Context, logger *zap.Logger) {
	r := router.New()
	logger.Info("app started successfully")

	server := &http.Server{
		Addr:    "8000",
		Handler: r,
	}

	idleConns := make(chan struct{})

	go func() {
		shutdownSig := make(chan os.Signal, 1)
		signal.Notify(shutdownSig, syscall.SIGTERM, syscall.SIGINT)
		<-shutdownSig

		logger.Info("received SIGTERM/SIGINT, starting shutdown")

		close(idleConns)
	}()

	if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error())
	}

	<-idleConns

	logger.Info("server shut down")
}
