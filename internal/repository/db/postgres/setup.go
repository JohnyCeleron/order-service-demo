package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"order-service/internal/configs"
	"order-service/internal/repository/model"
)

func SetupPostgres() (*gorm.DB, error) {
	cfg := configs.NewPostgresConfig()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		cfg.HostDB,
		cfg.UserAppDB,
		cfg.PasswordAppDB,
		cfg.NameAppDB,
		cfg.PortDB,
		cfg.SSLModeDB,
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
