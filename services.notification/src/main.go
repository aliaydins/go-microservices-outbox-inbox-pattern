package main

import (
	"github.com/aliaydins/oipattern/services.notification/src/config"
	"github.com/aliaydins/oipattern/services.notification/src/entity"
	notification "github.com/aliaydins/oipattern/services.notification/src/internal"
	worker "github.com/aliaydins/oipattern/services.notification/src/pkg/inbox_worker"
	"github.com/aliaydins/oipattern/shared/rabbitmq"
	"github.com/aliaydins/oipattern/shared/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("ERROR: Couldn't connect to the DB.")
	}
	db.Migrator().DropTable(&entity.Inbox{})
	db.AutoMigrate(&entity.Inbox{})

	r, err := rabbitmq.NewRabbitMQ(rabbitMQOptions)
	if err != nil {
		log.Printf("ERROR: Error occured when new RabbitMQ instance created")
		return
	}

	repo := notification.NewRepository(db)
	service := notification.NewService(repo)
	handler := notification.NewHandler(service)

	go worker.InboxWorker(r, repo, config.NotificationQueue, config.OrderExchange)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
