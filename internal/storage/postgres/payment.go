package storagePostgres

import (
	"gorm.io/gorm"

	modelDB "order-service/internal/models/db"
)

type PostgresPayment struct {
	PostgresStorage
}

func NewPayment(db *gorm.DB) *PostgresPayment {
	return &PostgresPayment{
		PostgresStorage{
			db: db,
		},
	}
}

func (storage *PostgresPayment) GetAll() ([]modelDB.Payment, error) {
	var payments []modelDB.Payment
	if err := storage.db.Find(&payments).Error; err != nil {
		return []modelDB.Payment{}, err
	}
	return payments, nil
}

func (storage *PostgresPayment) Get(id uint) (modelDB.Payment, error) {
	var payment modelDB.Payment
	if err := storage.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return modelDB.Payment{}, err
	}
	return payment, nil
}

func (storage *PostgresPayment) Add(payment modelDB.Payment) error {
	result := storage.db.Create(&payment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
