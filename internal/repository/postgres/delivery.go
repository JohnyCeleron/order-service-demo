package postgres

import (
	"gorm.io/gorm"

	"order-service/internal/domain"
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

func (storage *PostgresDelivery) GetAll() ([]domain.Delivery, error) {
	var deliveries []domain.Delivery
	if err := storage.db.Find(&deliveries).Error; err != nil {
		return []domain.Delivery{}, err
	}
	return deliveries, nil
}

func (storage *PostgresDelivery) Get(id uint) (domain.Delivery, error) {
	var delivery domain.Delivery
	if err := storage.db.Where("id = ?", id).First(&delivery).Error; err != nil {
		return domain.Delivery{}, err
	}
	return delivery, nil
}

func (storage *PostgresDelivery) Add(delivery domain.Delivery) error {
	result := storage.db.Create(&delivery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
