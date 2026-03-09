package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (s *Api) HandleDeleteTask(ctx context.Context, request api.HandleDeleteTaskRequestObject) (api.HandleDeleteTaskResponseObject, error) {
	err := s.usecase.DeleteTask(ctx, request.Body.UserUuid.String(), request.Body.TaskUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return api.HandleDeleteTask404JSONResponse{api.NotFoundJSONResponse{
				Code:    "Not found",
				Message: "either task_id or user_id not exist",
			}}, nil
		}
		return api.HandleDeleteTask500JSONResponse{api.InternalErrorJSONResponse{
			Code:    "server error",
			Message: "if was not your fault",
		}}, nil
	}
	log.Printf("user %v deleted task %v", request.Body.UserUuid.String(), request.Body.TaskUuid.String())
	return api.HandleDeleteTask204Response{}, nil
}
