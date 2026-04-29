package task

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/google/uuid"
)

type CreateTaskInput struct {
	Title string
	Description string
	Completed bool
	UserID string
}

type CreateTaskUseCase interface {
	Execute(ctx context.Context, input CreateTaskInput) error
}

type PostgreSQLCreateTaskUseCase struct {
	repo domain.TaskRepository
}

func NewPostgreSQLCreateTaskUseCase(repo domain.TaskRepository) *PostgreSQLCreateTaskUseCase {
	return &PostgreSQLCreateTaskUseCase{
		repo: repo,
	}
}

func (usecase *PostgreSQLCreateTaskUseCase) Execute(ctx context.Context, input CreateTaskInput) error {
	id := uuid.New().String()
	taskEntity := domain.NewTask(
		id, input.Title, input.Description, input.Completed, input.UserID,
	)
	return usecase.repo.Create(ctx, taskEntity)
}