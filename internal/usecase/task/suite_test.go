package task

import (
	"context"
	"testing"
	"todo-backend/internal/api"

	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite
	service *TaskService
	ctx     context.Context
	useCase api.TaskUseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()
	s.service = new(TaskService)
	s.useCase = NewTaskUseСase(s.service)
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
