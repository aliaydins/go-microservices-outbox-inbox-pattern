package entity

import (
	"gorm.io/gorm"
	"time"
)

type Inbox struct {
	gorm.Model
	ID            int `gorm:"primary_key;autoIncrement:true"`
	OrderID       int
	CustomerEmail string
	Processed     bool `gorm:"not null"`
	EventType     string
	CreatedAt     time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
