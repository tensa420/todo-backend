package task

import (
	"context"
)

func (u *TaskUsecase) DeleteTask(ctx context.Context, userID, taskID string) error {
	err := u.taskService.DeleteTask(ctx, userID, taskID)
	if err != nil {
		return err
	}
	return nil
}
