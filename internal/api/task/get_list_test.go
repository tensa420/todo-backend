package task

import (
	"todo-backend/internal/entity"
	api "todo-backend/pkg/openapi"

	"github.com/google/uuid"
)

func (s *ServerSuite) TestGetList() {
	exampleUserUUID := uuid.New()
	exampleEntityList := []entity.Task{
		entity.Task{
			Title:       "test",
			Description: "test",
			UserUUID:    exampleUserUUID,
			Status:      entity.TaskStatusNew,
		},
		entity.Task{
			Title:       "test",
			Description: "test",
			UserUUID:    exampleUserUUID,
			Status:      entity.TaskStatusNew,
		},
	}
	exampleRequest := &api.GetListOfTasksRequest{UserUUID: exampleUserUUID}
	wantList := &api.GetListOfTasksResponse{
		Tasks: []api.Task{
			api.Task{
				UserUUID:    api.NewOptUUID(exampleUserUUID),
				Title:       api.NewOptString("test"),
				Description: api.NewOptString("test"),
				Status:      api.NewOptTaskStatus("new"),
			},
			api.Task{
				UserUUID:    api.NewOptUUID(exampleUserUUID),
				Title:       api.NewOptString("test"),
				Description: api.NewOptString("test"),
				Status:      api.NewOptTaskStatus("new"),
			},
		},
	}
	s.useCase.On("GetListOfTasks", s.ctx, exampleUserUUID.String()).Return(exampleEntityList, nil)
	tasks, err := s.server.HandleGetListOfTasks(s.ctx, exampleRequest)

	s.NoError(err)
	s.Equal(wantList, tasks)
}
