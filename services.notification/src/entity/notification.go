package entity

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	ID          int `gorm:"primary_key;autoIncrement:true"`
	OrderID     int
	MessageSent bool
	CreatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
