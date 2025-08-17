package app

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func New() (*Application, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := createDatabase()
	if err != nil {
		return &Application{}, err
	}
	return &Application{
		db: db,
	}, nil
}
