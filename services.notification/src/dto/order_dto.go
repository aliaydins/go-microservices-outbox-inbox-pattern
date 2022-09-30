package dto

type OrderDTO struct {
	OrderID       int    `json:"order_id"`
	CustomerEmail string `json:"customer_email"`
	EventType     string `json:"event_type"`
}
