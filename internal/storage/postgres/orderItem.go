package storagePostgres

import (
	"gorm.io/gorm"

	modelDB "order-service/internal/models/db"
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

func (storage *PostgresOrderItem) GetAll() ([]modelDB.OrderItem, error) {
	var orderItem []modelDB.OrderItem
	if err := storage.db.Find(&orderItem).Error; err != nil {
		return []modelDB.OrderItem{}, err
	}
	return orderItem, nil
}

func (storage *PostgresOrderItem) Get(id uint) (modelDB.OrderItem, error) {
	var orderItem modelDB.OrderItem
	if err := storage.db.Where("id = ?", id).First(&orderItem).Error; err != nil {
		return modelDB.OrderItem{}, err
	}
	return orderItem, nil
}

func (storage *PostgresOrderItem) Add(orderItem modelDB.OrderItem) error {
	result := storage.db.Create(&orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
