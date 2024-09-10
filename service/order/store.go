package order

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

func (store *Store) GetOrdersByUserId(userID string) ([]types.Order, error) {
	query := `
		SELECT
			o.id,
			o.user_id,
			o.total,
			o.status,
			o.address,
			o.created_at,
			o.updated_at,
			COALESCE(json_agg(json_build_object('id', oi.id, 'product_id', oi.product_id, 'quantity', oi.quantity, 'price', oi.price, 'created_at', oi.created_at, 'updated_at', oi.updated_at)) FILTER (WHERE oi.id IS NOT NULL), '[]') AS order_items
		FROM
			orders o
		LEFT JOIN
			order_items oi ON o.id = oi.order_id
		WHERE o.user_id = $1
		GROUP BY
			o.id
	`

	rows, err := store.db.Query(query, userID)

	if err != nil {
		log.Printf("Error querying orders: %v", err)
		return nil, err
	}
	orders := make([]types.Order, 0)
	for rows.Next() {
		order, err := scanRowIntoOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return orders, nil
}

func scanRowIntoOrder(rows *sql.Rows) (*types.Order, error) {
	partialOrder := new(types.PartialOrder)
	err := rows.Scan(&partialOrder.ID, &partialOrder.UserID, &partialOrder.Total, &partialOrder.Status, &partialOrder.Address, &partialOrder.CreatedAt, &partialOrder.UpdatedAt, &partialOrder.OrderItems)
	if err != nil {
		log.Printf("Error scanning order: %v", err)
		return nil, err
	}

	order := types.Order{
		ID:         partialOrder.ID,
		UserID:     partialOrder.UserID,
		Total:      partialOrder.Total,
		Status:     partialOrder.Status,
		Address:    partialOrder.Address,
		CreatedAt:  partialOrder.CreatedAt,
		UpdatedAt:  partialOrder.UpdatedAt,
		OrderItems: make([]types.OrderItem, 0),
	}

	if err := json.Unmarshal([]byte(partialOrder.OrderItems), &order.OrderItems); err != nil {
		log.Printf("Error unmarshalling order items: %v", err)
		return nil, err
	}

	return &order, nil
}
