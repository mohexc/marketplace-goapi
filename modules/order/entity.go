package order

import (
	"marketplace-goapi/pkg/base"

	"github.com/google/uuid"
)

type OrderItem struct {
	base.Model
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}

type Order struct {
	base.Model
	UserID     uuid.UUID   `json:"user_id"`
	OrderItems []OrderItem `json:"order_items"`
}
