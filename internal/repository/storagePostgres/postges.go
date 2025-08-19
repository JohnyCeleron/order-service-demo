package storagePostgres

import "gorm.io/gorm"

type PostgresStorage struct {
	DB *gorm.DB
}
