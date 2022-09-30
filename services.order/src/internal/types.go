package order

import "github.com/aliaydins/oipattern/services.order/src/entity"

type OrderDTO struct {
	ID            int     `json:"id"`
	CustomerEmail string  `json:"customer_email"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
}

type OutboxDTO struct {
	ID            int     `json:"id"`
	IsSent        bool    `json:"is_sent"`
	EventType     string  `json:"event_type"`
	OrderID       int     `json:"order_id"`
	CustomerEmail string  `json:"customer_email"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
}

func OrderMapper(order *entity.Order) OrderDTO {
	dto := OrderDTO{
		ID:            order.ID,
		CustomerEmail: order.CustomerEmail,
		Name:          order.Name,
		Amount:        order.Amount,
	}
	return dto
}

func OutboxMapper(outbox *entity.Outbox) OutboxDTO {
	dto := OutboxDTO{
		ID:            outbox.ID,
		OrderID:       outbox.OrderID,
		IsSent:        outbox.IsSent,
		EventType:     outbox.EventType,
		CustomerEmail: outbox.CustomerEmail,
		Name:          outbox.Name,
		Amount:        outbox.Amount,
	}
	return dto
}
