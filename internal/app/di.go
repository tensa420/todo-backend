package app

import (
	"context"
	"log"
	"net"
	"os"
	"todo-backend/internal/api/task"
	"todo-backend/internal/client/task_service/handlers"
	task2 "todo-backend/internal/usecase/task"
	"todo-backend/pkg/task_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	server            *task.TaskServer
	usecase           *task2.TaskUseСase
	taskServiceClient *handlers.TaskServiceClient
}

func NewdiContainer() *diContainer {
	return &diContainer{}
}
func (d *diContainer) Server(ctx context.Context) *task.TaskServer {
	if d.server == nil {
		d.server = task.NewTaskServer(d.UseCase(ctx))
	}
	return d.server
}

func (d *diContainer) UseCase(ctx context.Context) *task2.TaskUseСase {
	if d.usecase == nil {
		d.usecase = task2.NewTaskUseСase(d.TaskServiceClient(ctx))
	}
	return d.usecase
}

func (d *diContainer) TaskServiceClient(ctx context.Context) *handlers.TaskServiceClient {
	if d.taskServiceClient != nil {
		return d.taskServiceClient
	}
	host := os.Getenv("TASK_SERVICE_CLIENT_HOST")
	port := os.Getenv("TASK_SERVICE_CLIENT_PORT")

	conn, err := initTaskServiceGRPCConnection(ctx, net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("error connecting to task service: %v", err)
	}

	generatedClient := task_service.NewTaskServiceClient(conn)
	d.taskServiceClient = handlers.NewTaskServiceClient(generatedClient)

	return d.taskServiceClient
}
func initTaskServiceGRPCConnection(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
