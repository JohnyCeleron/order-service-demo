include .env
export $(shell sed 's/=.*//' .env)

# Переменные (берутся из .env или имеют значения по умолчанию)
DB_USER ?= postgres
DB_PASSWORD ?= 
DB_NAME ?= postgres
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_SSLMODE ?= disable

APP_DB_USER?=app
APP_DB_PASSWORD?=secret
APP_DB_NAME?=wb_level0_db

swagger:
	swag init -g cmd/app/main.go -o ./docs --parseDependency --parseInternal

# Подключиться к БД через psql
db-connect:
	docker exec -it database psql -U $(APP_DB_USER) -d $(APP_DB_NAME)
run-producer:
	go run cmd/producer/main.go
run-app:
	go run cmd/app/main.go