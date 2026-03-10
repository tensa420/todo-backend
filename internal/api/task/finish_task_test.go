package task

import (
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestFinishTask() {
	exampleUserUUID, exampleTaskUUID := uuid.New(), uuid.New()
	exampleRequest := &api.FinishTaskRequest{TaskUUID: exampleTaskUUID, UserUUID: exampleUserUUID}

	s.useCase.On("FinishTask", s.ctx, exampleTaskUUID.String(), exampleUserUUID.String()).Return(nil)
	_, err := s.server.HandleFinishTask(s.ctx, exampleRequest)

	s.NoError(err)
}
