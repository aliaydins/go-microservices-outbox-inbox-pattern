package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID         int    `gorm:"primary_key;autoIncrement:true"`
	CustomerID int    `gorm:"not null; column:customer_id"`
	Name       string `gorm:"not null"`
	Amount     float64
	CreatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
