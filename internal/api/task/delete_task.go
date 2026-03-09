package task

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"
)

func (s *TaskServer) HandleDeleteTask(ctx context.Context, req *api.DeleteTaskRequest) (api.HandleDeleteTaskRes, error) {
	err := s.useCase.DeleteTask(ctx, req.UserUUID.String(), req.TaskUUID.String())
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return &api.HandleDeleteTaskNotFound{
				Code:    "NOT FOUND",
				Message: err.Error(),
			}, nil
		}
		return &api.HandleDeleteTaskInternalServerError{
			Code:    "INTERNAL ERROR",
			Message: err.Error(),
		}, nil
	}
	log.Printf("user %v deleted task %v", req.UserUUID.String(), req.TaskUUID.String())
	return &api.HandleDeleteTaskNoContent{}, nil
}
