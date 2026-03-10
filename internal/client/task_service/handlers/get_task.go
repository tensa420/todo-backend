package handlers

import (
	"context"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"github.com/google/uuid"
)

func (t *TaskServiceClient) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	task, err := t.generatedClient.GetTask(ctx, &task_service.GetTaskRequest{
		TaskUUID: taskUUID,
		UserUUID: userUUID,
	})
	if task == nil {
		return entity.Task{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.Task{}, err
	}
	return entity.Task{
		TaskUUID:    uuid.MustParse(task.Task.TaskUUID),
		Status:      ConvertProtoStatusToEntityStatus(task.Task.Status),
		Description: task.Task.Description,
		Title:       task.Task.Title,
	}, nil
}

func ConvertProtoStatusToEntityStatus(status task_service.TaskStatus) entity.TaskStatus {
	switch status {
	case task_service.TaskStatus_TASK_STATUS_NEW:
		return entity.TaskStatusNew
	default:
		return entity.TaskStatusFinished
	}
}
