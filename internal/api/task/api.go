package task

import "todo-backend/internal/usecase"

type Api struct {
	usecase usecase.ToDoBackendUseCase
}

func NewApi(usecase usecase.ToDoBackendUseCase) *Api {
	return &Api{
		usecase: usecase,
	}
}
