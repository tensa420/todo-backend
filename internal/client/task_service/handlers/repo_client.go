package handlers

import "todo-backend/pkg/task_service"

type TaskServiceClient struct {
	generatedClient task_service.TaskServiceClient
}

func NewTaskServiceClient(client task_service.TaskServiceClient) *TaskServiceClient {
	return &TaskServiceClient{
		generatedClient: client,
	}
}
