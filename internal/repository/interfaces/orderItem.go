package repository

import "order-service/internal/domain"

type OrderItemRepository interface {
	GetAll() ([]domain.OrderItem, error)
	Get(id uint) (domain.OrderItem, error)
	Add(orderItem domain.OrderItem) error
}
