package task

import (
	"context"
)

func (u *ToDoUsecase) FinishTask(ctx context.Context, taskID, userID string) error {
	err := u.repoClient.FinishTask(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
