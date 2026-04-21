package user

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Email string
	Password string
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) error
}

type PostgreSQLCreateUserUseCase struct {
	repo domain.UserRepository
}

func NewPostgreSQLCreateUserUseCase(
	repository domain.UserRepository,
) *PostgreSQLCreateUserUseCase {
	return &PostgreSQLCreateUserUseCase{
		repo: repository,
	}
}

func (usecase *PostgreSQLCreateUserUseCase) Execute(
	ctx context.Context, input CreateUserInput,
) error {
	id := uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInput, err := domain.NewUser(id, input.Email, string(hashPassword))
	if err != nil {
		return err
	}

	err = usecase.repo.Create(ctx, userInput)
	return err
}
