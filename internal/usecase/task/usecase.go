package task

import (
	"todo-backend/internal/usecase"
)

type TaskUseСase struct {
	taskService usecase.TaskService
}

func NewTaskUseСase(taskService usecase.TaskService) *TaskUseСase {
	return &TaskUseСase{taskService: taskService}
}
