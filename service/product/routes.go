package product

import (
	"go-api/types"
	"go-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
}

func (handler *Handler) GetProducts(response http.ResponseWriter, request *http.Request) {
	products, err := handler.store.GetProducts()
	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(response, http.StatusOK, products)
}

func (handler *Handler) CreateProduct(response http.ResponseWriter, request *http.Request) {
	var payload types.ProductPayload

	if err := utils.ParseJson(request, &payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	product, err := handler.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
	})

	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}

	var imagesArr []types.ProductImage

	if len(payload.Images) == 0 {
		productWithImages := types.ProductWithImages{
			Product: *product,
			Images:  []types.ProductImage{},
		}

		utils.WriteJson(response, http.StatusCreated, productWithImages)
		return
	}

	for _, image := range payload.Images {
		image, err := handler.store.CreateProductImage(types.ProductImage{
			ProductID: product.ID,
			ImageURL:  image.ImageURL,
		})

		imagesArr = append(imagesArr, *image)

		if err != nil {
			utils.WriteError(response, http.StatusInternalServerError, err)
			return
		}
	}

	productWithImages := types.ProductWithImages{
		Product: *product,
		Images:  imagesArr,
	}

	utils.WriteJson(response, http.StatusCreated, productWithImages)
}
