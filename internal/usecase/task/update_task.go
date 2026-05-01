package task

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
)

type UpdateTaskUseCase interface {
	Execute(ctx context.Context, taskID string, userID string, completed bool) error
}

type PostgreSQLUpdateTaskUseCase struct {
	repo domain.TaskRepository
}

func NewPostgreSQLUpdateTaskUseCase(repo domain.TaskRepository) *PostgreSQLUpdateTaskUseCase {
	return &PostgreSQLUpdateTaskUseCase{
		repo: repo,
	}
}

func (usecase *PostgreSQLUpdateTaskUseCase) Execute(
	ctx context.Context, taskID string, userID string, completed bool,
) error {
	return usecase.repo.Update(
		ctx, taskID, userID, completed,
	)
}
