package entity

import (
	"gorm.io/gorm"
)

type Inbox struct {
	gorm.Model
	ID        int  `gorm:"primary_key;autoIncrement:true"`
	Processed bool `gorm:"not null default true"`
	EventType string
	OrderID   int
}
