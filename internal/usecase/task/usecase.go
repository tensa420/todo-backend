package task

import (
	"todo-backend/internal/client/task_service"
)

type TaskUsecase struct {
	taskService task_service.TaskService
}

func NewTaskUsecase(taskService task_service.TaskService) *TaskUsecase {
	return &TaskUsecase{taskService: taskService}
}
