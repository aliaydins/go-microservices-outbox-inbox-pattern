package main

import (
	"fmt"
	"github.com/aliaydins/oipattern/services.order/src/config"
	"github.com/aliaydins/oipattern/services.order/src/entity"
	"github.com/aliaydins/oipattern/services.order/src/internal"
	worker "github.com/aliaydins/oipattern/services.order/src/pkg/outbox_worker"
	"github.com/aliaydins/oipattern/shared/rabbitmq"
	"github.com/aliaydins/oipattern/shared/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	config := config.LoadConfig(".")
	rabbitMQOptions := rabbitmq.RabbitMQOptions{
		URL:          config.GetRabbitMQURL(),
		RetryAttempt: 3,
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("ERROR: Couldn't connect to the DB.")
	}

	db.Migrator().DropTable(&entity.Order{}, &entity.Outbox{})
	db.AutoMigrate(&entity.Order{}, &entity.Outbox{})

	r, err := rabbitmq.NewRabbitMQ(rabbitMQOptions)
	if err != nil {
		log.Printf("ERROR: Error occured when new RabbitMQ instance created")
		return
	}

	err = r.CreateExchange(config.OrderExchange, "fanout", true, false)
	if err != nil {
		log.Printf("ERROR: Error occured when order-exchange created")
		return
	}

	repo := order.NewRepository(db)
	service := order.NewService(repo)
	handler := order.NewHandler(service)

	go worker.OutboxWorker(r, repo)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		fmt.Errorf("error %v", err.Error())
		panic("Couldn't start the HTTP server.")
	}

}
