package handlers

import (
	"context"
	"fmt"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *TaskServiceClient) DeleteTask(ctx context.Context, ID, userID string) error {
	_, err := t.generatedClient.DeleteTask(ctx, &task_service.DeleteTaskRequest{
		TaskID: ID,
		UserID: userID,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return fmt.Errorf("non grpc error %v", err)
		}
		switch st.Code() {
		case codes.NotFound:
			return entity.ErrNotFound
		default:
			return entity.ErrUnexpected
		}
	}
	log.Printf("delete task %v %v", ID, userID)
	return nil
}
