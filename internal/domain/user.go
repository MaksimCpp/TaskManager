package domain


type User struct {
	ID string
	Email string
	Password string
}

func NewUser(
	id string,
	email string,
	password string,
) *User {
	return &User{
		ID: id,
		Email: email, 
		Password: password,
	}
}
