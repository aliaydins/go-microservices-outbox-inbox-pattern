package worker

import (
	"encoding/json"
	"github.com/aliaydins/oipattern/services.notification/src/dto"
	"github.com/aliaydins/oipattern/services.notification/src/entity"
	notification "github.com/aliaydins/oipattern/services.notification/src/internal"
	"github.com/aliaydins/oipattern/shared/rabbitmq"
	"gorm.io/gorm"
	"log"
)

func InboxWorker(r *rabbitmq.RabbitMQ, repository *notification.Repository, queueName string, exchangeName string) {

	var orderDto dto.OrderDTO
	_, err := r.CreateQueue(queueName, true, false)
	if err != nil {
		log.Println("Error occured when deliveryQueue created")
		return
	}
	r.BindQueueWithExchange(queueName, "", exchangeName)
	r.CreateMessageChannel(queueName, "notification", true)

	go func() {
		for {
			message, err := r.ConsumeMessageChannel()
			if err != nil {
				log.Printf("Error occured when consuming message: %s\n", err.Error())
			}

			eventType := message.Headers["Key"]

			err = json.Unmarshal(message.Body, &orderDto)
			if err != nil {
				log.Printf("Can't unmarshal the byte array\n")
			}
			log.Printf("INFO: Message consumed with customer email -> %s and order id -> %d\n", orderDto.CustomerEmail, orderDto.OrderID)

			if eventType == "OrderCreated" {

				newInbox := entity.Inbox{
					OrderID:       orderDto.OrderID,
					Processed:     true,
					CustomerEmail: orderDto.CustomerEmail,
					EventType:     orderDto.EventType,
				}
				//search id
				searchedInbox, err := repository.GetInboxByOrderID(newInbox.OrderID)
				if err != gorm.ErrRecordNotFound {
					log.Printf("ERROR: Error occured when searched with order id -> %d\n", newInbox.OrderID)
				}

				if searchedInbox == nil {
					// you can do any operation you want here.
					log.Printf("INFO: Information Email send to %s successfully with order id -> %d\n", newInbox.CustomerEmail, newInbox.OrderID)
					// save the inbox processed = true
					err = repository.CreateInbox(&newInbox)
					if err != nil {
						log.Printf("ERROR: Error occured when created inbox")
					}
				} else {
					// repeated event
					log.Printf("ERROR: Information Email send to %s before with order id -> %d\n", newInbox.CustomerEmail, newInbox.OrderID)
				}

			}

		}
	}()

}
