package task

import (
	"context"
	"todo-backend/internal/entity"
)

func (u *TaskUsecase) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	taskID, err := u.taskService.CreateTask(ctx, task)
	if err != nil {
		return "", entity.ErrInternalServerError
	}
	return taskID, nil
}
