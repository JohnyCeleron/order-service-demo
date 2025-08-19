package storagePostgres

import (
	"gorm.io/gorm"

	"order-service/internal/repository/model"
)

type PostgresPayment struct {
	PostgresStorage
}

func NewPayment(db *gorm.DB) *PostgresPayment {
	return &PostgresPayment{
		PostgresStorage{
			DB: db,
		},
	}
}

func (storage *PostgresPayment) GetAll() ([]model.Payment, error) {
	var payments []model.Payment
	if err := storage.DB.Find(&payments).Error; err != nil {
		return []model.Payment{}, err
	}
	return payments, nil
}

func (storage *PostgresPayment) Get(id uint) (model.Payment, error) {
	var payment model.Payment
	if err := storage.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

func (storage *PostgresPayment) Add(payment model.Payment) error {
	result := storage.DB.Create(&payment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
