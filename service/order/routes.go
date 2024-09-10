package order

import (
	"go-api/types"
	"go-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("products/user/{id}", handler.GetOrdersByUserId).Methods("GET")
	router.HandleFunc("products", handler.CreateOrder).Methods("POST")
}

func (handler *Handler) GetOrdersByUserId(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	ID := params["id"]

	orders, err := handler.store.GetOrdersByUserId(ID)
	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(response, http.StatusOK, orders)
}

func (handler *Handler) CreateOrder(response http.ResponseWriter, request *http.Request) {
	var payload types.OrderPayload

	if err := utils.ParseJson(request, &payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	order, err := handler.store.CreateOrder(types.Order{
		UserID:  payload.UserID,
		Total:   payload.Total,
		Status:  payload.Status,
		Address: payload.Address,
	})

	for _, orderItem := range payload.OrderItems {
		_, err = handler.store.CreateOrderItem(types.OrderItem{
			OrderID:   order.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
			Price:     orderItem.UnitPrice,
		})

		if err != nil {
			utils.WriteError(response, http.StatusInternalServerError, err)
			return
		}
	}

	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(response, http.StatusCreated, order)
}
