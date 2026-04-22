package user

import (
	"context"
	"errors"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserInput struct {
	Email string
	Password string
}

type LoginUserOutput struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
}

func NewLoginUserOutput(id uuid.UUID, email string) *LoginUserOutput {
	return &LoginUserOutput{
		ID: id,
		Email: email,
	}
}

type LoginUserUseCase interface {
	Execute(ctx context.Context, input LoginUserInput) (*domain.User, error)
}

type PostgreSQLLoginUserUseCase struct {
	repo domain.UserRepository
}

func NewPostgreSQLLoginUserUseCase(
	repository domain.UserRepository,
) *PostgreSQLLoginUserUseCase {
	return &PostgreSQLLoginUserUseCase{
		repo: repository,
	}
}

func (usecase *PostgreSQLLoginUserUseCase) Execute(
	ctx context.Context, input LoginUserInput,
) (*domain.User, error) {
	user, err := usecase.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)
	if err != nil {
		return nil, errors.New("Invalid password.")
	}
	return user, nil
}
