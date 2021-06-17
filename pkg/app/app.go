package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/hendrikdelarey/cash-collection/pkg/router"
)

// Run starts the HTTP application server
func Run(ctx context.Context, logger *zap.Logger) {
	r := router.New()
	logger.Info("app started successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("errror creating api listen and serve: %e", err))
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
