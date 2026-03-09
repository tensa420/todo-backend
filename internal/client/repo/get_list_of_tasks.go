package repo

import (
	"context"
	"pet_api/internal/entity"
)

func (r *TaskServiceClient) GetListOfTasks(ctx context.Context) ([]entity.Task, error) {
	resp, err := r.generatedClient.GetListOfTasks(ctx, nil)
	if err != nil {
		return nil, err
	}
	tasks := make([]entity.Task, 0)
	for _, task := range resp.GetTasks() {
		tasks = append(task,entity.Task{
			ID: task.ID,
			Status: task.Status,
			Description: task.Description,
		})
	}
	return resp.Tasks., nil
}
