package task

import (
	"todo-backend/internal/usecase"
	"todo-backend/pkg/openapi"
)

type TaskServer struct {
	api.UnimplementedHandler
	useCase usecase.ToDoBackendUseCase
}

func NewApi(useCase usecase.ToDoBackendUseCase) *TaskServer {
	return &TaskServer{
		useCase: useCase,
	}
}
