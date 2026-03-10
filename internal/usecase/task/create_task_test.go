package task

import (
	"todo-backend/internal/entity"

	"github.com/google/uuid"
)

func (s *UseCaseSuite) TestCreateTask() {
	exampletask := entity.Task{
		UserUUID:    uuid.New(),
		Title:       "test",
		Description: "test",
		Status:      "NEW",
	}
	expectedUUID := uuid.New().String()
	s.service.On("CreateTask", s.ctx, exampletask).Return(expectedUUID, nil).Once()
	taskUUID, err := s.useCase.CreateTask(s.ctx, exampletask)

	s.NoError(err)
	s.Equal(taskUUID, expectedUUID)
}
