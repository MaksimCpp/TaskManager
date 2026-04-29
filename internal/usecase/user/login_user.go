package user

import (
	"context"
	"errors"

	"github.com/MaksimCpp/TaskManager/internal/domain"
	jwtservice "github.com/MaksimCpp/TaskManager/internal/service/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserInput struct {
	Email string
	Password string
}

type LoginUserOutput struct {
	Token string
}

type LoginUserUseCase interface {
	Execute(ctx context.Context, input LoginUserInput) (*LoginUserOutput, error)
}

type PostgreSQLLoginUserUseCase struct {
	repo domain.UserRepository
	jwt jwtservice.JWTService
}

func NewPostgreSQLLoginUserUseCase(
	repository domain.UserRepository,
	jwt jwtservice.JWTService,
) *PostgreSQLLoginUserUseCase {
	return &PostgreSQLLoginUserUseCase{
		repo: repository,
		jwt: jwt,
	}
}

func (usecase *PostgreSQLLoginUserUseCase) Execute(
	ctx context.Context, input LoginUserInput,
) (*LoginUserOutput, error) {
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

	token, err := usecase.jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("Invalid password.")
	}

	return &LoginUserOutput{
		Token: token,
	}, nil
}
