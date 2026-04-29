package domain

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetByUserID(ctx context.Context, userID string) ([]Task, error)
	Delete(ctx context.Context, taskID string, userID string) error
}