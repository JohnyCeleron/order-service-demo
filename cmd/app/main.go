package main

import (
	"log"

	"order-service/internal/app"
)

func main() {
	_, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

}
