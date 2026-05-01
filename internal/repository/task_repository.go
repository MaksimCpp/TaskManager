package repository

import (
	"context"
	"errors"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSQLTaskRepository struct {
	db *pgxpool.Pool
}

func NewPostgreSQLTaskRepository(db *pgxpool.Pool) *PostgreSQLTaskRepository {
	return &PostgreSQLTaskRepository{
		db: db,
	}
}

func (repo *PostgreSQLTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := `
		INSERT INTO taskschema.tasks (id, title, description, completed, user_id)
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err := repo.db.Exec(
		ctx,
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Completed,
		task.UserID,
	)
	return err
}

func (repo *PostgreSQLTaskRepository) GetByUserID(
	ctx context.Context, userID string,
) ([]domain.Task, error) {
	query := `
		SELECT id, title, description, completed, user_id
		FROM taskschema.tasks
		WHERE user_id = $1;
	`
	rows, err := repo.db.Query(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.UserID,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *PostgreSQLTaskRepository) Delete(
	ctx context.Context, taskID string, userID string,
) error {
	query := `
		DELETE FROM taskschema.tasks
		WHERE id = $1 AND user_id = $2;
	`
	_, err := repo.db.Exec(
		ctx,
		query,
		taskID,
		userID,
	)
	return err
}

func (repo *PostgreSQLTaskRepository) Update(
	ctx context.Context, taskID string, userID string, completed bool,
) error {
	query := `
		UPDATE taskschema.tasks
		SET completed = $1, updated_at = NOW()
		WHERE id = $2 AND user_id = $3;
	`
	result, err := repo.db.Exec(
		ctx,
		query,
		completed,
		taskID,
		userID,
	)

	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
        return errors.New("Task or user not found")
    }
	return nil
}
