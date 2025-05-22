package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/acgtubio/eydsh-auth/internal/auth"
	"github.com/acgtubio/eydsh-auth/routes"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		println("Some error.")
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	w io.Writer,
	args []string,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// TODO: Implement provider for auth service
	authService, err := auth.DefaultAuthService()
	if err != nil {
		return err
	}

	server := routes.NewServer(authService)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	go func() {
		log.Printf("Listening on port 8080...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error starting server: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}
