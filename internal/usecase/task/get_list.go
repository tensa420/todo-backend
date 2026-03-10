package task

import (
	"context"
	"todo-backend/internal/entity"
)

func (u *TaskUseСase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	tasks, err := u.taskService.GetListOfTasks(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
