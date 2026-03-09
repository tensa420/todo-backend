package task

import (
	"context"
	"errors"
	"log"

	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (a *Api) HandleGetListOfTasks(ctx context.Context, request api.HandleGetListOfTasksRequestObject) (api.HandleGetListOfTasksResponseObject, error) {
	tasks, err := a.usecase.GetListOfTasks(ctx, request.Body.UserUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return api.HandleGetListOfTasks404JSONResponse{api.NotFoundJSONResponse{
				Code:    "Not found",
				Message: "user_id not exist",
			}}, nil
		}
		return api.HandleGetListOfTasks404JSONResponse{api.NotFoundJSONResponse{
			Code:    "server error",
			Message: "if was not your fault",
		}}, nil
	}
	log.Printf("tasks`s list %v of user %v ", tasks, request.Body.UserUuid.String())
	return api.HandleGetListOfTasks200JSONResponse{Tasks: convert(tasks)}, nil
}

func convert(tasks []entity.Task) []api.Task {
	var converted []api.Task
	for _, task := range tasks {
		status := convertTaskStatus(task.Status)
		converted = append(converted, api.Task{
			Description: &task.Description,
			Title:       &task.Title,
			Status:      &status,
		})
	}
	return converted
}
func convertTaskStatus(status entity.TaskStatus) api.TaskStatus {
	switch status {
	case "new":
		temp := api.New
		return api.TaskStatus{Status: &temp}
	default:
		temp := api.Finished
		return api.TaskStatus{Status: &temp}
	}
}
