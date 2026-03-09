package task_service

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"
)

func (r *TaskServiceClient) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	resp, err := r.generatedClient.CreateTask(ctx, &task_service.CreateTaskRequest{
		Description: task.Description,
		Status:      task_service.Status_STATUS_NEW,
		Title:       task.Title,
		UserID:      task.UserID.String(),
	})
	if err != nil {
		return "", entity.ErrInternalServerError
	}
	log.Printf("create task %v %v", task.UserID, task.Title)
	return resp.ID, nil
}
