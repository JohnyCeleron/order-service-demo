package main

import (
	"log"

	"github.com/joho/godotenv"

	"order-service/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	application.Run()
}
