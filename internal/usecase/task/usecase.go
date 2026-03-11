package task

import (
	"todo-backend/internal/usecase"
)

type TaskUseСase struct {
	taskService usecase.TaskServiceTypes
}

func NewTaskUseСase(taskService usecase.TaskServiceTypes) *TaskUseСase {
	return &TaskUseСase{taskService: taskService}
}
