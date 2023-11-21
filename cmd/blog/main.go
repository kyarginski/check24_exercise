// Blog service for CHECK24.
//
// # Description of the REST API of the service for working with blogs and comments.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http, https
// Host: localhost
// Version: 1.0.0
//
// swagger:meta
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"check24/internal/api"
	"check24/internal/config"
	"check24/internal/database"
	"check24/internal/lib/logger/sl"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.MustLoad("blog")
	log := sl.SetupLogger(cfg.Env)
	log.Info(
		"starting blog server",
		slog.String("env", cfg.Env),
		slog.String("version", cfg.Version),
	)

	if err := run(log, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}
}

func run(log *slog.Logger, cfg *config.Config) error {
	log.Debug("starting blog server")

	log.Debug("starting db connect ", "connect", cfg.DBConnect)
	storage, err := database.New(cfg.DBConnect)
	if err != nil {
		return err
	}
	defer storage.Close()

	log.Debug("db connected successfully")

	router := mux.NewRouter()

	router.Use(api.ValidateTokenMiddleware)

	// Define API endpoints for creating and reading blog entries.
	router.HandleFunc("/api/blog/entries", api.GetBlogEntries(storage)).Methods("GET")
	router.HandleFunc("/api/blog/entries", api.CreateBlogEntry(storage)).Methods("POST")
	router.HandleFunc("/api/blog/entries/{entry_id}", api.GetBlogEntry(storage)).Methods("GET")

	/*
		router.HandleFunc("/api/blog/entries/{entry_id}", api.UpdateBlogEntry(db)).Methods("PUT")

		// Define API endpoints for creating and reading comments.
		router.HandleFunc("/api/blog/entries/{entry_id}/comments", api.CreateComment).Methods("POST")
		router.HandleFunc("/api/blog/entries/{entry_id}/comments", api.GetCommentsForEntry).Methods("GET")

		// Define API endpoints for deleting blog entries.
		router.HandleFunc("/api/blog/entries/{entry_id}", api.DeleteBlogEntry).Methods("DELETE")
	*/
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	cfgAddress := fmt.Sprintf(":%d", cfg.Port)

	srv := &http.Server{
		Addr:    cfgAddress,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error("failed to start server", "error", err)
			}
		}
	}()

	log.Info("started blog server", "port", cfgAddress)

	<-done
	log.Info("stopping blog server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", "error", err)

		return err
	}

	log.Info("blog server stopped")

	return nil
}
