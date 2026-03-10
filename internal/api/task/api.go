package task

import (
	"todo-backend/internal/usecase"
	"todo-backend/pkg/openapi"
)

type TaskServer struct {
	useCase usecase.TaskUseCase
	api.UnimplementedHandler
}

func NewTaskServer(useCase usecase.TaskUseCase) *TaskServer {
	return &TaskServer{
		useCase: useCase,
	}
}
