package task

import (
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestCreateTask() {
	taskUUID := uuid.New()
	userUUID := uuid.New()
	expectedUUID := &api.CreateTaskResponse{TaskUUID: taskUUID}

	wantTask := &api.Task{
		TaskUUID:    api.NewOptNilUUID(taskUUID),
		UserUUID:    api.NewOptUUID(userUUID),
		Title:       api.NewOptString("test"),
		Description: api.NewOptString("test"),
		Status:      api.NewOptTaskStatus("new"),
	}
	exampleTask := entity.Task{
		TaskUUID:    taskUUID,
		UserUUID:    userUUID,
		Description: "test",
		Title:       "test",
		Status:      "new",
	}
	s.useCase.On("CreateTask", s.ctx, exampleTask).Return(expectedUUID.TaskUUID.String(), nil).Once()

	result, err := s.server.HandleCreateTask(s.ctx, wantTask)

	s.NoError(err)
	s.Equal(expectedUUID, result)
}
