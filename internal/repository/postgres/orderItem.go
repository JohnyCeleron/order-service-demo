package postgres

import (
	"gorm.io/gorm"

	"order-service/internal/repository/model"
)

type PostgresOrderItem struct {
	PostgresStorage
}

func NewOrderItem(db *gorm.DB) *PostgresPayment {
	return &PostgresPayment{
		PostgresStorage{
			db: db,
		},
	}
}

func (storage *PostgresOrderItem) GetAll() ([]model.OrderItem, error) {
	var orderItem []model.OrderItem
	if err := storage.db.Find(&orderItem).Error; err != nil {
		return []model.OrderItem{}, err
	}
	return orderItem, nil
}

func (storage *PostgresOrderItem) Get(id uint) (model.OrderItem, error) {
	var orderItem model.OrderItem
	if err := storage.db.Where("id = ?", id).First(&orderItem).Error; err != nil {
		return model.OrderItem{}, err
	}
	return orderItem, nil
}

func (storage *PostgresOrderItem) Add(orderItem model.OrderItem) error {
	result := storage.db.Create(&orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
