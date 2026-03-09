package task

import (
	"pet_api/internal/client/repo"
)

type ToDoBackendUsecase struct {
	taskService repo.RepoClient
}
