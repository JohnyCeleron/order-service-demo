package postgres

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	domainOrder "order-service/internal/domain/order"
	"order-service/internal/repository/converter"
	"order-service/internal/repository/model"
)

type PostgresStorage struct {
	DB *gorm.DB
}

type PostgresOrder struct {
	PostgresStorage
}

func New() (*PostgresOrder, error) {
	db, err := SetupPostgres()
	if err != nil {
		return &PostgresOrder{}, err
	}
	return &PostgresOrder{
		PostgresStorage{
			DB: db,
		},
	}, nil
}

func (storage *PostgresOrder) GetAll() ([]domainOrder.Order, error) {
	var orders []model.Order
	if err := storage.DB.
		Joins("Delivery").
		Joins("Payment").
		Preload("Items").
		Find(&orders).Error; err != nil {
		return []domainOrder.Order{}, err
	}
	domainOrders := make([]domainOrder.Order, len(orders))
	for i, order := range orders {
		domainOrders[i] = converter.OrderModelDBToDomain(order)
	}
	return domainOrders, nil
}

func (storage *PostgresOrder) Get(id string) (domainOrder.Order, error) {
	var order model.Order
	if err := storage.DB.
		Joins("Delivery").
		Joins("Payment").
		Preload("Items").
		Where("orders.id = ?", id).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domainOrder.Order{}, fmt.Errorf("order with id %s not found", id)
		}
		return domainOrder.Order{}, fmt.Errorf("failed to get order: %v", err)
	}
	return converter.OrderModelDBToDomain(order), nil
}

func (storage *PostgresOrder) Add(order domainOrder.Order) error {
	modelOrder := converter.OrderDomainToModelDB(order)
	result := storage.DB.Create(&modelOrder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (storage *PostgresOrder) Close() error {
	sqlDB, err := storage.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
