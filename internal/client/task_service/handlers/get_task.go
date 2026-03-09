package handlers

import (
	"context"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"github.com/google/uuid"
)

func (r *TaskServiceClient) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	task, err := r.generatedClient.GetTask(ctx, &task_service.GetTaskRequest{
		TaskID: taskUUID,
		UserID: userUUID,
	})
	if task == nil {
		return entity.Task{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.Task{}, entity.ErrInternalServerError
	}
	return entity.Task{
		ID:          uuid.MustParse(task.ID),
		Status:      ConvertProtoStatusToEntityStatus(task.Status.String()),
		Description: task.Description,
		Title:       task.Title,
	}, nil
}

func ConvertProtoStatusToEntityStatus(status string) entity.TaskStatus {
	switch status {
	case "NEW":
		return entity.TaskStatusNew
	default:
		return entity.TaskStatusDone
	}
}
