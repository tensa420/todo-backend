package task

import (
	task_service2 "todo-backend/internal/client/task_service"
)

type ToDoBackendUsecase struct {
	taskService task_service2.TaskService
}

func NewToDoBackendUsecase(taskService task_service2.TaskService) *ToDoBackendUsecase {
	return &ToDoBackendUsecase{taskService: taskService}
}
