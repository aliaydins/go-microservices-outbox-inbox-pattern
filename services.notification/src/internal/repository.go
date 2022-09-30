package notification

import (
	"github.com/aliaydins/oipattern/services.notification/src/entity"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateInbox(order *entity.Inbox) error {
	return r.db.Model(&entity.Inbox{}).Create(&order).Error
}

func (r *Repository) GetInboxByOrderID(orderId int) (*entity.Inbox, error) {
	inbox := new(entity.Inbox)
	err := r.db.Where("order_id = ?", orderId).First(&inbox).Error
	if err != nil {
		return nil, err
	}
	return inbox, nil
}

func (r *Repository) GetInboxList() ([]entity.Inbox, error) {
	var list []entity.Inbox
	err := r.db.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
