// wf-dba: main.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"data-rest/pkg"

	db "github.com/Peter-Bird/Flash-DB"
)

func init() {
	appName := filepath.Base(os.Args[0])
	log.SetPrefix("[" + appName + "] ")
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	// Create a context that listens for the interrupt signal
	ctx, stop := createInterruptContext(ctx)
	defer stop()

	// Initialize the application (configuration, dependencies, etc.)
	server, err := initApp(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize application: %w", err)
	}

	// Channel to capture server errors
	errCh := make(chan error, 1)

	// Run the server in a goroutine
	go func() {
		log.Printf("Server running on port %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("server error: %w", err)
		}
	}()

	// Wait for context cancellation or server errors
	select {
	case <-ctx.Done():
		log.Println("Received shutdown signal")
	case err := <-errCh:
		log.Printf("Server error: %v\n", err)
		return err
	}

	// Gracefully shut down the server
	return shutdownServer(ctx, server)
}

func initApp(ctx context.Context) (*http.Server, error) {
	// Load configuration
	cfg := pkg.LoadConfig()

	// Initialize dependencies
	repo := db.NewFlashDB()
	service := pkg.NewService(repo)
	handler := pkg.NewHandler(service)
	router := pkg.NewRouter(handler)

	// Create server
	server := pkg.NewServer(cfg.Port, router)
	return server, nil
}

func createInterruptContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, os.Interrupt)
}

func shutdownServer(ctx context.Context, server *http.Server) error {
	log.Println("Shutting down server...")

	// Create a deadline for the shutdown process
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	log.Println("Server exited gracefully")
	return nil
}
