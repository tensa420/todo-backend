package handlers

import (
	"context"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"github.com/google/uuid"
)

func (t *TaskServiceClient) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	task, err := t.generatedClient.GetTask(ctx, &task_service.GetTaskRequest{
		TaskID: taskUUID,
		UserID: userUUID,
	})
	if task == nil {
		return entity.Task{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.Task{}, err
	}
	return entity.Task{
		TaskUUID:    uuid.MustParse(task.ID),
		Status:      ConvertProtoStatusToEntityStatus(task.Status),
		Description: task.Description,
		Title:       task.Title,
	}, nil
}

func ConvertProtoStatusToEntityStatus(status task_service.Status) entity.TaskStatus {
	switch status {
	case task_service.Status_STATUS_NEW:
		return entity.TaskStatusNew
	default:
		return entity.TaskStatusDone
	}
}
