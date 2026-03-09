package task

import (
	"context"
	"todo-backend/internal/entity"
)

func (u *ToDoBackendUsecase) GetTask(ctx context.Context, userID, taskID string) (entity.Task, error) {
	task, err := u.taskService.GetTask(ctx, userID, taskID)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}
