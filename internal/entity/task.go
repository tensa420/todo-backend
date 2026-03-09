package entity

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	UserID      uuid.UUID  `json:"user_id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

type TaskStatus string

const (
	TaskStatusDone TaskStatus = "finished"
	TaskStatusNew  TaskStatus = "new"
)
