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

func (t *TaskServiceClient) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	_, err := t.generatedClient.FinishTask(ctx, &task_service.FinishTaskRequest{
		UserUUID: userUUID,
		TaskUUID: taskUUID,
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
			return entity.ErrUnexpected
		}
	}
	log.Printf("finish task %v %v", taskUUID, userUUID)
	return nil
}
