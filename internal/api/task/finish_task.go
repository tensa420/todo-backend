package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (s *TaskServer) HandleFinishTask(ctx context.Context, req *api.FinishTaskRequest) (api.HandleFinishTaskRes, error) {
	err := s.useCase.FinishTask(ctx, req.TaskUUID.String(), req.UserUUID.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return &api.HandleFinishTaskNotFound{
				Code:    "NOT FOUND",
				Message: err.Error(),
			}, nil
		}
		return &api.HandleFinishTaskInternalServerError{
			Code:    "INTERNAL ERROR",
			Message: err.Error(),
		}, nil
	}

	log.Printf("user %v finished task %v", req.TaskUUID.String(), req.UserUUID.String())
	return &api.HandleFinishTaskNoContent{}, nil
}
