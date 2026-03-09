package task_service

import (
	"context"
	"fmt"
	"log"
	"todo-backend/internal/entity"
	"todo-backend/pkg/task_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *TaskServiceClient) DeleteTask(ctx context.Context, ID, userID string) error {
	_, err := r.generatedClient.DeleteTask(ctx, &task_service.DeleteTaskRequest{
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
		case codes.Internal:
			return entity.ErrInternalServerError
		default:
			return entity.ErrUnexpected
		}
	}
	log.Printf("delete task %v %v", ID, userID)
	return nil
}
