package repository

import (
	"order-service/internal/domain"
)

type PaymentRepository interface {
	GetAll() ([]domain.Payment, error)
	Get(id uint) (domain.Payment, error)
	Add(payment domain.Payment) error
}
