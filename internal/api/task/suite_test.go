package task

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	ctx     context.Context
	useCase *TaskUseCase
	server  *TaskServer
}

func (s *ServerSuite) SetupTest() {
	s.ctx = context.Background()
	s.useCase = new(TaskUseCase)
	s.server = NewTaskServer(s.useCase)
}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
