package usecase

import (
	"context"
	"todo-backend/internal/entity"
)

type ToDoBackendUseCase interface {
	DeleteTask(ctx context.Context, userID, taskID string) error
	GetTask(ctx context.Context, userID, taskID string) (entity.Task, error)
	GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error)
	FinishTask(ctx context.Context, taskID, userID string) error
	CreateTask(ctx context.Context, task entity.Task) (string, error)
}
