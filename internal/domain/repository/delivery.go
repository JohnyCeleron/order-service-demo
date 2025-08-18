package repository

import modelDB "order-service/internal/models/db"

type DeliveryRepository interface {
	GetAll() ([]modelDB.Delivery, error)
	Get(id uint) (modelDB.Delivery, error)
	Add(delivery modelDB.Delivery) error
}
