package storagePostgres

import "gorm.io/gorm"

type PostgresStorage struct {
	db *gorm.DB
}
