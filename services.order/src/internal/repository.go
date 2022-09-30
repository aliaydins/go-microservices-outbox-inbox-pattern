package order

import (
	"github.com/aliaydins/oipattern/services.order/src/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateOrder(order *entity.Order) error {
	return r.db.Model(&entity.Order{}).Create(&order).Error
}

func (r *Repository) CreateOutbox(outbox *entity.Outbox) error {
	return r.db.Model(&entity.Outbox{}).Create(&outbox).Error
}

func (r *Repository) GetOutboxList() ([]entity.Outbox, error) {
	var list []entity.Outbox
	err := r.db.Where("is_sent = ?", false).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *Repository) UpdateStatus(outbox *entity.Outbox) error {
	return r.db.Model(&outbox).Where("id = ?", outbox.ID).Updates(entity.Outbox{IsSent: true}).Error
}
