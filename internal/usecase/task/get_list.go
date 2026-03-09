package task

import (
	"context"
	"todo-backend/internal/entity"
)

func (u *ToDoBackendUsecase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	tasks, err := u.GetListOfTasks(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
