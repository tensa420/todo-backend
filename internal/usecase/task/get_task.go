package task

import (
	"context"
	"pet_api/internal/entity"
)

func (u *ToDoUsecase)GetTask(ctx context.Context, userID,taskID string) (entity.Task, error) {
	task,err := u.
}
