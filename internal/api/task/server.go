package task

import (
	"todo-backend/pkg/openapi"
)

type TaskServer struct {
	useCase TaskUseCaseTypes
	openapi.UnimplementedHandler
}

func NewTaskServer(useCase TaskUseCaseTypes) *TaskServer {
	return &TaskServer{
		useCase: useCase,
	}
}
