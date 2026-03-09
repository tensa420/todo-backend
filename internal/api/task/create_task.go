package task

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *Api) HandleCreateTask(ctx context.Context, request api.HandleCreateTaskRequestObject) (api.HandleCreateTaskResponseObject, error) {
	id, err := s.usecase.CreateTask(ctx, entity.Task{
		Description: request.Body.Description,
		Title:       request.Body.Title,
		UserID:      request.Body.UserUuid,
		Status:      entity.TaskStatusNew,
	})
	if err != nil {
		log.Printf("failed to create task: %v", err)
		return api.HandleCreateTask500JSONResponse{api.InternalErrorJSONResponse{
			Code:    "INTERNAL_ERROR",
			Message: "failed to create task",
		}}, nil
	}

	log.Printf("user %v created task: %v", request.Body.UserUuid.String(), id)
	return api.HandleCreateTask201JSONResponse{
		TaskUuid: uuid.MustParse(id),
	}, nil
}
