package task

import "github.com/google/uuid"

func (s *UseCaseSuite) TestDeleteTask() {
	exampleTaskUUID, exampleUserUUID := uuid.New().String(), uuid.New().String()
	s.service.On("DeleteTask", s.ctx, exampleTaskUUID, exampleUserUUID).Return(nil).Once()
	err := s.useCase.DeleteTask(s.ctx, exampleTaskUUID, exampleUserUUID)
	s.NoError(err)
}
