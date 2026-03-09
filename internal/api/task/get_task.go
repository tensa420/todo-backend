package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (a *Api) HandleGetTask(ctx context.Context, request api.HandleGetTaskRequestObject) (api.HandleGetTaskResponseObject, error) {
	task, err := a.usecase.GetTask(ctx, request.Body.UserUuid.String(), request.Body.TaskUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return api.HandleGetTask404JSONResponse{api.NotFoundJSONResponse{
				Code:    "Not found",
				Message: "either task_uuid or user_uuid not exist",
			}}, nil
		}
		return api.HandleGetTask404JSONResponse{api.NotFoundJSONResponse{
			Code:    "server error",
			Message: "if was not your fault",
		}}, nil
	}
	log.Printf("user`s %v got task %v", request.Body.UserUuid, request.Body.TaskUuid)
	return api.HandleGetTask200JSONResponse{Description: task.Description, Title: task.Title, Status: convertTaskStatus(task.Status)}, nil
}
