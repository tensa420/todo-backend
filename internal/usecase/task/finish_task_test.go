package task

import "github.com/google/uuid"

func (s *UseCaseSuite) TestFinishTask() {
	exampleTaskUUID, exampleUserUUID := uuid.NewString(), uuid.NewString()
	s.service.On("FinishTask", s.ctx, exampleTaskUUID, exampleUserUUID).Return(nil).Once()
	err := s.useCase.FinishTask(s.ctx, exampleTaskUUID, exampleUserUUID)
	s.NoError(err)
}
