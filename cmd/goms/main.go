package main

import (
	"context"
	"fmt"
	"github.com/deepPublicGit/go-microservice/internal/handlers"
	"log/slog"
	"net/http"
	"os"
	"time"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/deepPublicGit/go-microservice/internal/log"
)

func main() {
	// Logger configuration
	logger := log.New(
		log.WithLevel(os.Getenv("LOG_LEVEL")),
		log.WithSource(),
	)

	if err := run(logger); err != nil {
		logger.ErrorContext(context.Background(), "an error occurred", slog.String("error", err.Error()))
		os.Exit(1)
	}

	router := handlers.RouterHandler(logger)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.ErrorContext(context.Background(), "an error occurred", slog.String("error", err.Error()))
	}
}

func run(logger *slog.Logger) error {
	ctx := context.Background()

	_, err := maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		logger.DebugContext(ctx, fmt.Sprintf(s, i...))
	}))
	if err != nil {
		return fmt.Errorf("setting max procs: %w", err)
	}

	logger.InfoContext(ctx, "Hello world!", slog.String("location", "world"))

	return nil
}
