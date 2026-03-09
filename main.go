package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	task2 "todo-backend/internal/api/task"
	"todo-backend/internal/client/task_service/handlers"
	"todo-backend/internal/usecase/task"
	"todo-backend/pkg/openapi"
	"todo-backend/pkg/task_service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host, port := os.Getenv("TASK_SERVICE_CLIENT_HOST"), os.Getenv("TASK_SERVICE_CLIENT_PORT")
	address := net.JoinHostPort(host, port)
	conn, err := initTaskServiceGRPCConnection(ctx, address)
	if err != nil {
		log.Printf("failed to make connection")
	}
	generatedClient := task_service.NewTaskServiceClient(conn)
	taskServiceClient := handlers.NewTaskServiceClient(generatedClient)
	usecase := task.NewToDoBackendUsecase(taskServiceClient)
	api := task2.NewApi(usecase)
	openapi.Handler(api)
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
