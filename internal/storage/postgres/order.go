package storagePostgres

import (
	"gorm.io/gorm"

	modelDB "order-service/internal/models/db"
)

type PostgresOrder struct {
	PostgresStorage
}

func NewOrder(db *gorm.DB) *PostgresOrder {
	return &PostgresOrder{
		PostgresStorage{
			db: db,
		},
	}
}

func (storage *PostgresOrder) GetAll() ([]modelDB.Order, error) {
	var orders []modelDB.Order
	if err := storage.db.Find(&orders).Error; err != nil {
		return []modelDB.Order{}, err
	}
	return orders, nil
}

func (storage *PostgresOrder) Get(id string) (modelDB.Order, error) {
	var order modelDB.Order
	if err := storage.db.Where("id = ?", id).First(&order).Error; err != nil {
		return modelDB.Order{}, err
	}
	return order, nil
}

func (storage *PostgresOrder) Add(order modelDB.Order) error {
	result := storage.db.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
