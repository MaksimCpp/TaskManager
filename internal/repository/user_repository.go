package repository

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSQLUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgreSQLUserRepository(db *pgxpool.Pool) *PostgreSQLUserRepository {
	return &PostgreSQLUserRepository{
		db: db,
	}
}

func (repo *PostgreSQLUserRepository) Create(
	ctx context.Context, user *domain.User,
) error {
	query := `
		INSERT INTO taskschema.users (id, email, password, created_at)
		VALUES ($1, $2, $3, NOW());
	`
	_, err := repo.db.Exec(
		ctx,
		query,
		user.ID,
		user.Email,
		user.Password,
	)
	return err
}

func (repo *PostgreSQLUserRepository) GetByEmail(
	ctx context.Context, email string,
) (*domain.User, error) {
	query := `
		SELECT id, email, password
		FROM taskschema.users
		WHERE email = $1;
	`
	var user domain.User
	err := repo.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgreSQLUserRepository) GetByID(
	ctx context.Context, id uuid.UUID,
) (*domain.User, error) {
	query := `
		SELECT id, email, password
		FROM taskschema.users
		WHERE id = $1;
	`
	var user domain.User
	err := repo.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
