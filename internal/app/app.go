package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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

func (a *App) initDI(ctx context.Context) error {
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
		log.Printf("failed to listen: %v", err)
	}
	log.Printf("server started")
	err = srv.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
	return nil
}
func (a *App) Run(ctx context.Context) error {
	err := a.RunHTTPServer(ctx)
	if err != nil {
		return err
	}
	return nil
}
