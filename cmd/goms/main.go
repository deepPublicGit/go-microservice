package main

import (
	"context"
	"fmt"
	"github.com/deepPublicGit/go-microservice/internal/handlers"
	"github.com/gorilla/mux"
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
	/*	homeHandler := handlers.NewJobHandler(logger)
		servMux := http.NewServeMux()
		servMux.Handle("/", homeHandler)

		server := &http.Server{
			Addr:    ":8080",
			Companies: servMux,
		}
		err := server.ListenAndServe()
		if err != nil {
			logger.ErrorContext(context.Background(), "an error occurred", slog.String("error", err.Error()))
		}*/
	router := mux.NewRouter()
	router.Handle("/", handlers.NewHomeHandler(logger))

	jobsRouter := router.PathPrefix("/jobs").Subrouter()
	jobsRouter.Handle("/", handlers.NewJobHandler(logger))
	jobsRouter.Handle("/{id}", handlers.NewJobHandler(logger))
	jobsRouter.Handle("/{id:[0-9]+}/company", handlers.NewJobHandler(logger))

	companyHandler := handlers.NewCompanies(logger)
	companiesRouter := router.PathPrefix("/companies").Subrouter()
	companiesRouter.HandleFunc("/", companyHandler.GetCompanies)
	companiesRouter.Handle("/{id}", handlers.NewCompanies(logger))
	companiesRouter.Handle("/{company}/jobs", handlers.NewCompanies(logger))

	// Get, Get Batch, Post, Post Batch, Later Get Pagination, Patch single.
	// Post Company, Get Company,Companies,Put,Patch
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
