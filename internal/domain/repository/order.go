package repository

import modelDB "order-service/internal/models/db"

type OrderRepository interface {
	GetAll() ([]modelDB.Order, error)
	Get(id string) (modelDB.Order, error)
	Add(order modelDB.Order) error
}
