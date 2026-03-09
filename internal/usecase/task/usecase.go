package task

import (
	task_service2 "todo-backend/internal/client/task_service"
)

type ToDoBackendUsecase struct {
	taskService task_service2.TaskService
}
