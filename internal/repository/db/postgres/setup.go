package postgres

import (
	"fmt"
	"order-service/internal/repository/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupPostgres() (*gorm.DB, error) {
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
		return &gorm.DB{}, fmt.Errorf("failed to connect to database:", err)
	}
	db.AutoMigrate(&model.Order{}, &model.Delivery{}, &model.OrderItem{}, &model.Payment{})
	return db, nil
}
