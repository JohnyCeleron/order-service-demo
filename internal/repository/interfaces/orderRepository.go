package repository

import "order-service/internal/domain"

type OrderRepository interface {
	GetAll() ([]domain.Order, error)
	Get(id string) (domain.Order, error)
	Add(order domain.Order) error
}
