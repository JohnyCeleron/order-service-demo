package repository

import modelDB "order-service/internal/models/db"

type PaymentRepository interface {
	GetAll() ([]modelDB.Payment, error)
	Get(id uint) (modelDB.Payment, error)
	Add(payment modelDB.Payment) error
}
