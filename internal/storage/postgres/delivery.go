package storagePostgres

import (
	"gorm.io/gorm"

	modelDB "order-service/internal/models/db"
)

type PostgresDelivery struct {
	PostgresStorage
}

func NewDelivery(db *gorm.DB) *PostgresDelivery {
	return &PostgresDelivery{
		PostgresStorage{
			db: db,
		},
	}
}

func (storage *PostgresDelivery) GetAll() ([]modelDB.Delivery, error) {
	var deliveries []modelDB.Delivery
	if err := storage.db.Find(&deliveries).Error; err != nil {
		return []modelDB.Delivery{}, err
	}
	return deliveries, nil
}

func (storage *PostgresDelivery) Get(id uint) (modelDB.Delivery, error) {
	var delivery modelDB.Delivery
	if err := storage.db.Where("id = ?", id).First(&delivery).Error; err != nil {
		return modelDB.Delivery{}, err
	}
	return delivery, nil
}

func (storage *PostgresDelivery) Add(delivery modelDB.Delivery) error {
	result := storage.db.Create(&delivery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
