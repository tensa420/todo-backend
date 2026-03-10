package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	task2 "todo-backend/internal/api/task"
	"todo-backend/internal/client/task_service/handlers"
	"todo-backend/internal/usecase/task"
	api "todo-backend/pkg/openapi"
	"todo-backend/pkg/task_service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
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
	grpcClient := task_service.NewTaskServiceClient(conn)
	taskServiceClient := handlers.NewTaskServiceClient(grpcClient)
	useCase := task.NewToDoBackendUsecase(taskServiceClient)
	apii := task2.NewApi(useCase)
	handler, err := api.NewServer(apii)
	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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
