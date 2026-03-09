package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (a *TaskServer) HandleGetListOfTasks(ctx context.Context, req *api.GetListOfTasksRequest) (api.HandleGetListOfTasksRes, error) {
	tasks, err := a.useCase.GetListOfTasks(ctx, req.UserUUID.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return &api.HandleGetListOfTasksNotFound{
				Code:    "NOT FOUND",
				Message: err.Error(),
			}, nil
		}
		return &api.HandleGetListOfTasksInternalServerError{
			Code:    "INTERNAL ERROR",
			Message: err.Error(),
		}, nil
	}
	log.Printf("tasks`s list %v of user %v ", tasks, req.UserUUID.String())
	return &api.GetListOfTasksResponse{Tasks: convert(tasks)}, nil
}

func convert(tasks []entity.Task) []api.Task {
	var converted []api.Task
	for _, task := range tasks {
		status := convertTaskStatus(task.Status)
		converted = append(converted, api.Task{
			Description: api.NewOptString(task.Description),
			Title:       api.NewOptString(task.Title),
			UserUUID:    api.NewOptUUID(task.UserUUID),
			Status:      api.NewOptTaskStatus(status),
		})
	}
	return converted
}
func convertTaskStatus(status entity.TaskStatus) api.TaskStatus {
	switch status {
	case entity.TaskStatusNew:
		return api.TaskStatusNew
	default:
		return api.TaskStatusFinished
	}
}
