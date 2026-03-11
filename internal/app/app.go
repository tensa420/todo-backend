package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	api "todo-backend/pkg/openapi"
)

type App struct {
	diContainer *diContainer
	httpServer  *api.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	err := a.initDI(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewdiContainer()
	return nil
}
func (a *App) RunHTTPServer(ctx context.Context) error {
	httpServer := a.diContainer.Server(ctx)

	handler, err := api.NewServer(httpServer)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    net.JoinHostPort(os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT")),
		Handler: handler,
	}

	lis, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("server started on %s", srv.Addr)
		if err = srv.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err = srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
func (a *App) Run(ctx context.Context) error {
	err := a.RunHTTPServer(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (a *App) gracefulShutdown(ctx context.Context, srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Println("Received signal:", sig)
	case <-ctx.Done():
		log.Println("Context cancelled")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
