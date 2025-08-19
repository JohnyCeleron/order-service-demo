package postgres

import (
	"gorm.io/gorm"

	"order-service/internal/domain"
	"order-service/internal/repository/converter"
	"order-service/internal/repository/model"
)

type PostgresOrder struct {
	PostgresStorage
}

func NewOrder(db *gorm.DB) *PostgresOrder {
	return &PostgresOrder{
		PostgresStorage{
			db: db,
		},
	}
}

func (storage *PostgresOrder) GetAll() ([]domain.Order, error) {
	var orders []model.Order
	if err := storage.db.Find(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	domainOrders := make([]domain.Order, len(orders))
	for i, order := range orders {
		domainOrders[i] = converter.OrderModelDBToDomain(order)
	}
	return domainOrders, nil
}

func (storage *PostgresOrder) Get(id string) (domain.Order, error) {
	var order model.Order
	if err := storage.db.Where("id = ?", id).First(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return converter.OrderModelDBToDomain(order), nil
}

func (storage *PostgresOrder) Add(order domain.Order) error {
	modelOrder := converter.OrderDomainToModelDB(order)
	result := storage.db.Create(&modelOrder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
