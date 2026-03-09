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

func (r *TaskServiceClient) FinishTask(ctx context.Context, taskID, userID string) error {
	_, err := r.generatedClient.FinishTask(ctx, &task_service.FinishTaskRequest{
		UserID: userID,
		TaskID: taskID,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return fmt.Errorf("non grpc error %v", err)
		}
		switch st.Code() {
		case codes.NotFound:
			return entity.ErrNotFound
		default:
			return entity.ErrInternalServerError
		}
	}
	log.Printf("finish task %v %v", taskID, userID)
	return nil
}
