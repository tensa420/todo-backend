package task

import (
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestDeleteTask() {
	exampleUserUUID, exampleTaskUUID := uuid.New(), uuid.New()
	exampleRequest := &api.DeleteTaskRequest{TaskUUID: exampleTaskUUID, UserUUID: exampleUserUUID}

	s.useCase.On("DeleteTask", s.ctx, exampleUserUUID.String(), exampleTaskUUID.String()).Return(nil)

	_, err := s.server.HandleDeleteTask(s.ctx, exampleRequest)

	s.NoError(err)
}
