package types

import (
	"go-api/enums"
)

type RegisterPayload struct {
	Name     string     `json:"name" validate:"required"`
	Email    string     `json:"email" validate:"required,email"`
	Password string     `json:"password" validate:"required"`
	Role     enums.Role `json:"role" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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

type ProductImagesPayload struct {
	ProductID string `json:"product_id" validate:"required"`
	ImageURL  string `json:"image_url" validate:"required"`
}

type ProductPayload struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description" validate:"required"`
	Price       int                    `json:"price" validate:"required"`
	Images      []ProductImagesPayload `json:"images" validate:"required"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductImage struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(Product) (*Product, error)
}
