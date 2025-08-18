package repository

import modelDB "order-service/internal/models/db"

type OrderItemRepository interface {
	GetAll() ([]modelDB.OrderItem, error)
	Get(id uint) (modelDB.OrderItem, error)
	Add(orderItem modelDB.OrderItem) error
}
