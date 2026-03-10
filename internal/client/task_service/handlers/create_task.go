package handlers

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"
)

func (t *TaskServiceClient) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	resp, err := t.generatedClient.CreateTask(ctx, &task_service.CreateTaskRequest{
		Task: taskEntityToProto(task),
	})
	if err != nil {
		return "", err
	}
	log.Printf("create task %v %v", task.UserUUID, task.Title)
	return resp.TaskUUID, nil
}

func taskEntityToProto(task entity.Task) *task_service.Task {
	return &task_service.Task{
		Title:       task.Title,
		Description: task.Description,
		UserUUID:    task.UserUUID.String(),
		Status:      task_service.TaskStatus_TASK_STATUS_NEW,
	}
}
