package types

import "go-api/enums"

type SignUpPayload struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     enums.Role `json:"role"`
}

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      enums.Role `json:"role"`
	Password  string     `json:"password"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(User) error
}
