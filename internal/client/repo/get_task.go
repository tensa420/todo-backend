package repo

import (
	"context"
	"pet_api/internal/entity"
	"pet_api/pkg/task_service"

	"github.com/google/uuid"
)

func (r *TaskServiceClient) GetTask(ctx context.Context, ID string) (entity.Task, error) {
	task, err := r.generatedClient.GetTask(ctx, &task_service.GetTaskRequest{
		ID: ID,
	})
	if err != nil {
		return entity.Task{}, err
	}
	return entity.Task{
		ID:          uuid.MustParse(task.ID),
		Status:      ConvertProtoStatusToEntityStatus(task.Status.String()),
		Description: task.Description,
	}, nil
}

func ConvertProtoStatusToEntityStatus(status string) entity.Status {
	switch status {
	case "CREATED":
		return entity.StatusCreated
	default:
		return entity.StatusDone
	}
}
