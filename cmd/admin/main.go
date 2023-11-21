// Admin service for CHECK24.
//
// # Description of the REST API of the service for working with auth.
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
	"check24/internal/lib/logger/sl"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	cfg := config.MustLoad("admin")
	log := sl.SetupLogger(cfg.Env)
	log.Info(
		"starting admin server",
		slog.String("env", cfg.Env),
		slog.String("version", cfg.Version),
	)

	if err := run(log, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}
}

func run(log *slog.Logger, cfg *config.Config) error {
	log.Debug("starting admin server")

	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")

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
				log.Error("failed to start admin server", "error", err)
			}
		}
	}()

	log.Info("started admin server", "port", cfgAddress)

	<-done
	log.Info("stopping admin server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", "error", err)

		return err
	}

	log.Info("blog server stopped")

	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO DEBUG Simulate user authentication (in a real-world scenario, this would involve user verification)
	username := "root"
	password := "password"

	// Check user credentials (replace with actual user authentication logic).
	if r.FormValue("username") != username || r.FormValue("password") != password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid credentials")
		return
	}

	// Create a JWT token
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: r.FormValue("username"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(api.SecretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error creating token")
		return
	}

	// Send the token as a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}
