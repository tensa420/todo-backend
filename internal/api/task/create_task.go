package task

import (
	"context"
	"log"
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *TaskServer) HandleCreateTask(ctx context.Context, req *api.Task) (api.HandleCreateTaskRes, error) {
	taskUUID, err := s.useCase.CreateTask(ctx, TaskToEntityFromAPI(req))
	if err != nil {
		log.Printf("failed to create task: %v", err)
		return &api.HandleCreateTaskInternalServerError{
			Code:    "INTERNAL ERROR",
			Message: err.Error(),
		}, nil
	}

	log.Printf("user %v created task: %v", req.UserUUID, taskUUID)
	return &api.CreateTaskResponse{TaskUUID: uuid.MustParse(taskUUID)}, nil
}

func TaskToEntityFromAPI(req *api.Task) entity.Task {
	return entity.Task{
		Description: req.Description.Value,
		Title:       req.Title.Value,
		UserUUID:    req.UserUUID.Value,
		TaskUUID:    req.TaskUUID.Value,
		Status:      entity.TaskStatusNew,
	}
}
