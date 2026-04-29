package task

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
)

type DeleteTaskInput struct {
	TaskID string
	UserID string
}

type DeleteTaskUseCase interface {
	Execute(ctx context.Context, input DeleteTaskInput) error
}

type PostgreSQLDeleteTaskUseCase struct {
	repo domain.TaskRepository
}

func NewPostgreSQLDeleteTaskUseCase(repo domain.TaskRepository) *PostgreSQLDeleteTaskUseCase {
	return &PostgreSQLDeleteTaskUseCase{
		repo: repo,
	}
}

func (usecase *PostgreSQLDeleteTaskUseCase) Execute(
	ctx context.Context, input DeleteTaskInput,
) error {
	return usecase.repo.Delete(ctx, input.TaskID, input.UserID)
}
