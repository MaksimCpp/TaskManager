package task

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
)

type GetTasksOutput struct {
	Title string
	Completed bool
}

type GetTasksUseCase interface {
	Execute(ctx context.Context, userID string) ([]GetTasksOutput, error)
}

type PostgreSQLGetTasksUseCase struct {
	repo domain.TaskRepository
}

func NewPostgreSQLGetTasksUseCase(repo domain.TaskRepository) *PostgreSQLGetTasksUseCase {
	return &PostgreSQLGetTasksUseCase{
		repo: repo,
	}
}

func (usecase *PostgreSQLGetTasksUseCase) Execute(
	ctx context.Context, userID string,
) ([]GetTasksOutput, error) {
	tasks, err := usecase.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var tasksOutput []GetTasksOutput
	for _, task := range tasks {
		t := GetTasksOutput{
			Title: task.Title,
			Completed: task.Completed,
		}
		tasksOutput = append(tasksOutput, t)
	}
	return tasksOutput, nil
}