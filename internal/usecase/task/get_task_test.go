package task

import (
	"todo-backend/internal/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetTask() {
	wantTask := entity.Task{
		Title:       "test",
		Description: "test",
		UserUUID:    uuid.New(),
		TaskUUID:    uuid.New(),
		Status:      mock.Anything,
	}
	exampleTaskUUID, exampleUserUUID := uuid.New().String(), uuid.New().String()
	s.service.On("GetTask", s.ctx, exampleTaskUUID, exampleUserUUID).Return(wantTask, nil).Once()
	task, err := s.service.GetTask(s.ctx, exampleTaskUUID, exampleUserUUID)
	s.NoError(err)
	s.Equal(wantTask, task)
}
