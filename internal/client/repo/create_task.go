package repo

import (
	"context"
	"pet_api/internal/entity"
	"pet_api/pkg/task_service"
)

func (r *TaskServiceClient) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	resp, err := r.generatedClient.CreateTask(ctx, &task_service.CreateTaskRequest{
		Description: task.Description,
		Status:      task_service.Status_STATUS_CREATED,
	})
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
