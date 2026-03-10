package handlers

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"github.com/google/uuid"
)

func (t *TaskServiceClient) GetListOfTasks(ctx context.Context, userID string) ([]entity.Task, error) {
	resp, err := t.generatedClient.GetListOfTasks(ctx, &task_service.GetListOfTasksRequest{
		UserID: userID,
	})
	if resp == nil || len(resp.Tasks) == 0 {
		return []entity.Task{}, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0, len(resp.Tasks))
	for _, task := range resp.GetTasks() {
		tasks = append(tasks, entity.Task{
			Status:      ConvertProtoStatusToEntityStatus(task.Status),
			Description: task.Description,
			Title:       task.Title,
			TaskUUID:    uuid.MustParse(task.ID),
		})
	}

	log.Printf("get list tasks %v", userID)
	return tasks, nil
}
