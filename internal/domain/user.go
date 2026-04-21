package domain

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
	Email string
	Password string
}

func NewUser(
	id uuid.UUID,
	email string,
	password string,
) (*User, error) {
	return &User{
		ID: id,
		Email: email, 
		Password: password,
	}, nil
}
