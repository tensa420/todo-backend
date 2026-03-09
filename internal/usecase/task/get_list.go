package task

import (
	"context"
	"pet_api/internal/entity"
)

func (u *ToDoBackendUsecase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	return nil, nil
}
