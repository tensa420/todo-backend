package task

import (
	"todo-backend/internal/api"
	"todo-backend/pkg/openapi"
)

type TaskServer struct {
	useCase api.TaskUseCase
	openapi.UnimplementedHandler
}

func NewTaskServer(useCase api.TaskUseCase) *TaskServer {
	return &TaskServer{
		useCase: useCase,
	}
}
