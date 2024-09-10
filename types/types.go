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
	ImageURL string `json:"image_url" validate:"required"`
}

type ProductPayload struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description" validate:"required"`
	Price       float32                `json:"price" validate:"required"`
	Images      []ProductImagesPayload `json:"images"`
}

type Product struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	Images      []ProductImage `json:"images,omitempty"`
}

type PartialProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Images      string  `json:"images"`
}

type ProductWithImages struct {
	Product Product        `json:"product"`
	Images  []ProductImage `json:"images"`
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
	GetProductById(id string) (*Product, error)
	CreateProduct(Product) (*Product, error)
	CreateProductImage(ProductImage) (*ProductImage, error)
	UpdateProduct(Product) (*Product, error)
	DeleteProduct(id string) error
}

type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	Total      float32     `json:"total"`
	Status     string      `json:"status"`
	Address    string      `json:"address"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
	OrderItems []OrderItem `json:"order_items,omitempty"`
}

type PartialOrder struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	Total      float32 `json:"total"`
	Status     string  `json:"status"`
	Address    string  `json:"address"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	OrderItems string  `json:"order_items"`
}

type OrderItem struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type OrderPayload struct {
	UserID     string             `json:"user_id" validate:"required"`
	Total      float32            `json:"total" validate:"required"`
	Status     string             `json:"status" validate:"required"`
	Address    string             `json:"address" validate:"required"`
	OrderItems []OrderItemPayload `json:"order_items_ids"`
}

type OrderItemPayload struct {
	OrderID   string  `json:"order_id" validate:"required"`
	ProductID string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float32 `json:"unit_price" validate:"required"`
}

type OrderStore interface {
	CreateOrder(Order) (*Order, error)
	GetOrdersByUserId(userId string) ([]Order, error)
	CreateOrderItem(OrderItem) (*OrderItem, error)
}
