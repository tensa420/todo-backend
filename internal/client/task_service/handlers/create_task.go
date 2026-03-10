package handlers

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"
)

func (t *TaskServiceClient) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	resp, err := t.generatedClient.CreateTask(ctx, TaskEntityToProto(task))
	if err != nil {
		return "", entity.ErrInternalServerError
	}
	log.Printf("create task %v %v", task.UserUUID, task.Title)
	return resp.ID, nil
}

func TaskEntityToProto(task entity.Task) *task_service.CreateTaskRequest {
	return &task_service.CreateTaskRequest{
		Title:       task.Title,
		Description: task.Description,
		UserID:      task.UserUUID.String(),
		Status:      task_service.Status_STATUS_NEW,
	}
}
