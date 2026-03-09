package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (s *Api) HandleFinishTask(ctx context.Context, request api.HandleFinishTaskRequestObject) (api.HandleFinishTaskResponseObject, error) {
	err := s.usecase.FinishTask(ctx, request.Body.TaskUuid.String(), request.Body.UserUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return api.HandleFinishTask404JSONResponse{api.NotFoundJSONResponse{
				Code:    "Not found",
				Message: "either task_id or user_id not exist",
			}}, nil
		}
		return api.HandleFinishTask500JSONResponse{api.InternalErrorJSONResponse{
			Code:    "server error",
			Message: "if was not your fault",
		}}, nil
	}

	log.Printf("user %v finished task %v", request.Body.UserUuid.String(), request.Body.TaskUuid.String())
	return api.HandleFinishTask204Response{}, nil
}
