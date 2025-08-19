package repository

import (
	"order-service/internal/domain"
)

type DeliveryRepository interface {
	GetAll() ([]domain.Delivery, error)
	Get(id uint) (domain.Delivery, error)
	Add(delivery domain.Delivery) error
}
