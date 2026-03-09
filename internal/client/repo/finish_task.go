package repo

import (
	"context"
	"errors"
	"pet_api/internal/entity"
	"pet_api/pkg/task_service"

	"github.com/jackc/pgx/v5"
)

func (r *TaskServiceClient) FinishTask(ctx context.Context, ID, userID string) error {
	_, err := r.generatedClient.FinishTask(ctx, &task_service.FinishTaskRequest{
		TaskID: ID,
		UserID: userID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.ErrNotFound
		}
		return err
	}
	return nil
}
