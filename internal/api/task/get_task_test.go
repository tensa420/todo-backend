package task

import (
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestGetTask() {
	exampleUserUUID, exampleTaskUUID := uuid.New(), uuid.New()
	exampleTask := entity.Task{
		TaskUUID:    exampleTaskUUID,
		UserUUID:    exampleUserUUID,
		Description: "test",
		Title:       "test",
		Status:      entity.TaskStatusNew,
	}
	wantTask := &api.Task{
		TaskUUID:    api.NewOptNilUUID(exampleTask.TaskUUID),
		UserUUID:    api.NewOptUUID(exampleUserUUID),
		Description: api.NewOptString("test"),
		Title:       api.NewOptString("test"),
		Status:      api.NewOptTaskStatus(api.TaskStatusNew),
	}
	exampleRequest := &api.GetTaskRequest{TaskUUID: exampleTask.TaskUUID, UserUUID: exampleUserUUID}
	s.useCase.On("GetTask", s.ctx, exampleUserUUID.String(), exampleTaskUUID.String()).Return(exampleTask, nil)

	task, err := s.server.HandleGetTask(s.ctx, exampleRequest)

	s.NoError(err)
	s.Equal(wantTask, task)
}
