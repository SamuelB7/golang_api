package product

import (
	"bytes"
	"encoding/json"
	"go-api/types"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	handler := NewHandler(productStore)

	t.Run("Should fail if the product payload is invalid", func(t *testing.T) {
		payload := types.ProductPayload{
			Name:        "test",
			Description: "test",
			Price:       0,
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.CreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should create a new product", func(t *testing.T) {

		payload := types.ProductPayload{
			Name:        "test",
			Description: "test",
			Price:       10,
			Images:      []types.ProductImagesPayload{},
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.CreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
		}

		log.Printf("Response: %v", rr.Body.String())
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) GetProductById(id string) (*types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) CreateProduct(product types.Product) (*types.Product, error) {
	product.ID = "abc-123"
	return &product, nil
}

func (m *mockProductStore) UpdateProduct(product types.Product) (*types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) DeleteProduct(id string) error {
	return nil
}

func (m *mockProductStore) CreateProductImage(productImage types.ProductImage) (*types.ProductImage, error) {
	return nil, nil
}
