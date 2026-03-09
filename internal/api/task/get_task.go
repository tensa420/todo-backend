package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (a *TaskServer) HandleGetTask(ctx context.Context, req *api.GetTaskRequest) (api.HandleGetTaskRes, error) {
	task, err := a.useCase.GetTask(ctx, req.UserUUID.String(), req.TaskUUID.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return &api.HandleGetTaskNotFound{
				Code:    "NOT_FOUND",
				Message: err.Error(),
			}, nil
		}
		return &api.HandleGetTaskInternalServerError{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		}, nil
	}
	log.Printf("user`s %v got task %v", req.UserUUID, req.TaskUUID)
	return &api.Task{
		Description: api.NewOptString(task.Description),
		Title:       api.NewOptString(task.Title),
		Status:      api.NewOptTaskStatus(convertTaskStatus(task.Status)),
		UserUUID:    api.NewOptUUID(task.UserUUID),
		TaskUUID:    api.NewOptNilUUID(task.TaskUUID),
	}, nil
}
