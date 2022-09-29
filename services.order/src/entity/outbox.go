package entity

import (
	"gorm.io/gorm"
	"time"
)

type Outbox struct {
	gorm.Model
	ID         int  `gorm:"primary_key;autoIncrement:true"`
	IsSent     bool `gorm:"not null default false"`
	EventType  string
	OrderID    int
	CustomerID int
	Name       string
	Amount     float64
	CreatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
