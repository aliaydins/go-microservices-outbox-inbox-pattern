package worker

import (
	"encoding/json"
	"github.com/aliaydins/oipattern/services.order/src/config"
	"github.com/aliaydins/oipattern/services.order/src/event"
	order "github.com/aliaydins/oipattern/services.order/src/internal"
	"github.com/aliaydins/oipattern/shared/rabbitmq"
	"github.com/robfig/cron/v3"
	"log"
)

func OutboxWorker(r *rabbitmq.RabbitMQ, repo *order.Repository) {
	c := cron.New()
	c.AddFunc("@every 1m", func() {

		outboxList, err := repo.GetOutboxList()
		if err != nil {
			log.Println("ERROR: Error occured when getting outbox list from database")
			return
		}

		if len(outboxList) != 0 {
			for _, e := range outboxList {
				payload, _ := json.Marshal(event.OrderCreated{OrderID: e.OrderID, CustomerEmail: e.CustomerEmail, EventType: e.EventType})

				err := r.Publish(payload, config.AppConfig.OrderExchange, e.EventType)
				if err != nil {
					log.Printf("ERROR: Error occured when published event wiht orderID  -> %d\n", e.ID)
					return
				}

				err = repo.UpdateStatus(&e)
				if err != nil {
					log.Printf("ERROR: Error occured when event status updated with orderID -> %d\n", e.ID)
					return
				}
				log.Printf("INFO: Event published with orderID -> %d\n", e.ID)
			}

		}

	})
	c.Start()
}
