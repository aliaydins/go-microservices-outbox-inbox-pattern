package notification

import "github.com/aliaydins/oipattern/services.notification/src/entity"

type InboxDTO struct {
	ID            int    `json:"id"`
	OrderID       int    `json:"order_id"`
	Processed     bool   `json:"processed"`
	EventType     string `json:"event_type"`
	CustomerEmail string `json:"customer_email"`
}

func InboxMapper(inbox *entity.Inbox) InboxDTO {
	dto := InboxDTO{
		ID:            inbox.ID,
		OrderID:       inbox.OrderID,
		Processed:     inbox.Processed,
		EventType:     inbox.EventType,
		CustomerEmail: inbox.CustomerEmail,
	}
	return dto
}
