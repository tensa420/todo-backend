package task

import (
	"todo-backend/internal/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetList() {
	exampleListOfTask := []entity.Task{
		entity.Task{
			Title:       "title",
			Description: "description",
			TaskUUID:    uuid.New(),
			UserUUID:    uuid.New(),
			Status:      mock.Anything,
		},
		entity.Task{
			Title:       "title1",
			Description: "description1",
			TaskUUID:    uuid.New(),
			UserUUID:    uuid.New(),
			Status:      mock.Anything,
		},
	}
	exampleUserUUID := uuid.New().String()
	s.service.On("GetListOfTasks", s.ctx, exampleUserUUID).Return(exampleListOfTask, nil).Once()
	tasks, err := s.useCase.GetListOfTasks(s.ctx, exampleUserUUID)
	
	s.NoError(err)
	s.Equal(exampleListOfTask, tasks)
}
