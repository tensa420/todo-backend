package repo

import "pet_api/pkg/task_service"

type TaskServiceClient struct {
	generatedClient task_service.TaskServiceClient
}

func NewRepoClient(client task_service.TaskServiceClient) *TaskServiceClient {
	return &TaskServiceClient{
		generatedClient: client,
	}
}
