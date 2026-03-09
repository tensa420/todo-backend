package task

import (
	"context"
)

func (u *ToDoBackendUsecase) FinishTask(ctx context.Context, taskID, userID string) error {
	err := u.taskService.FinishTask(ctx, taskID, userID)
	if err != nil {
		return err
	}
	return nil
}
