package main

import (
	"log"

	"order-service/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

}
