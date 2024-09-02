package product

import (
	"database/sql"
	"go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetProducts() ([]types.Product, error) {
	rows, err := store.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (store *Store) CreateProduct(product types.Product) (*types.Product, error) {
	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price"
	err := store.db.QueryRow(query, product.Name, product.Description, product.Price).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (store *Store) CreateProductImage(productImage types.ProductImage) error {
	_, err := store.db.Exec("INSERT INTO product_images (product_id, image_url) VALUES ($1, $2)", productImage.ProductID, productImage.ImageURL)
	if err != nil {
		return err
	}

	return nil
}
