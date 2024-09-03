package product

import (
	"database/sql"
	"encoding/json"
	"go-api/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetProducts() ([]types.Product, error) {
	rows, err := store.db.Query(`
		SELECT 
			p.id,
			p.name,
			p.description,
			p.price,
			p.created_at,
			p.updated_at,
			COALESCE(json_agg(json_build_object('id', pi.id, 'image_url', pi.image_url, 'created_at', pi.created_at, 'updated_at', pi.updated_at)) FILTER (WHERE pi.id IS NOT NULL), '[]') AS images
		FROM 
			products p
		LEFT JOIN 
			product_images pi ON p.id = pi.product_id
		GROUP BY 
			p.id;
	`)

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
	partialProduct := new(types.PartialProduct)
	err := rows.Scan(&partialProduct.ID, &partialProduct.Name, &partialProduct.Description, &partialProduct.Price, &partialProduct.CreatedAt, &partialProduct.UpdatedAt, &partialProduct.Images)
	if err != nil {
		log.Printf("Error scanning product: %v", err)
		return nil, err
	}

	var images []types.ProductImage

	err = json.Unmarshal([]byte(partialProduct.Images), &images)
	if err != nil {
		log.Printf("Error unmarshalling images: %v", err)
		return nil, err
	}

	product := &types.Product{
		ID:          partialProduct.ID,
		Name:        partialProduct.Name,
		Description: partialProduct.Description,
		Price:       partialProduct.Price,
		CreatedAt:   partialProduct.CreatedAt,
		UpdatedAt:   partialProduct.UpdatedAt,
		Images:      images,
	}

	return product, nil
}

func (store *Store) CreateProduct(product types.Product) (*types.Product, error) {
	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price, created_at, updated_at"
	err := store.db.QueryRow(query, product.Name, product.Description, product.Price).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, nil)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, err
	}

	return &product, nil
}

func (store *Store) CreateProductImage(productImage types.ProductImage) (*types.ProductImage, error) {
	query := "INSERT INTO product_images (product_id, image_url) VALUES ($1, $2) RETURNING id, product_id, image_url, created_at, updated_at"
	err := store.db.QueryRow(query, productImage.ProductID, productImage.ImageURL).Scan(&productImage.ID, &productImage.ProductID, &productImage.ImageURL, &productImage.CreatedAt, &productImage.UpdatedAt)
	if err != nil {
		log.Printf("Error creating product image: %v", err)
		return nil, err
	}

	return &productImage, nil
}
