package order

import "github.com/aliaydins/oipattern/services.order/src/entity"

type OrderDTO struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}

func mapper(order *entity.Order) OrderDTO {
	dto := OrderDTO{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Name:       order.Name,
		Amount:     order.Amount,
	}
	return dto
}
