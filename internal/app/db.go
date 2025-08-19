package app

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"order-service/internal/repository/model"
	"order-service/internal/repository/storagePostgres"
)

func NewPostgresOrder() (*storagePostgres.PostgresOrder, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return &storagePostgres.PostgresOrder{}, fmt.Errorf("failed to connect to database:", err)
	}
	db.AutoMigrate(&model.Order{}, &model.Delivery{}, &model.OrderItem{}, &model.Payment{})
	return &storagePostgres.PostgresOrder{
		PostgresStorage: storagePostgres.PostgresStorage{
			DB: db,
		},
	}, nil
}
