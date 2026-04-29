package user

import (
	"context"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Email string
	Password string
}

type RegisterUserUseCase interface {
	Execute(ctx context.Context, input RegisterUserInput) error
}

type PostgreSQLRegisterUserUseCase struct {
	repo domain.UserRepository
}

func NewPostgreSQLRegisterUserUseCase(
	repository domain.UserRepository,
) *PostgreSQLRegisterUserUseCase {
	return &PostgreSQLRegisterUserUseCase{
		repo: repository,
	}
}

func (usecase *PostgreSQLRegisterUserUseCase) Execute(
	ctx context.Context, input RegisterUserInput,
) error {
	id := uuid.New().String()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInput := domain.NewUser(id, input.Email, string(hashPassword))
	err = usecase.repo.Create(ctx, userInput)
	return err
}
