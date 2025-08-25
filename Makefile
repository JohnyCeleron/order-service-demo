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

MIGRATIONS_DIR ?= ./db/migrations

# Формируем строку подключения
DSN := "user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=$(DB_SSLMODE)"

.PHONY: migrate-up migrate-down migrate-status migrate-create migrate-reset

# Применить все новые миграции
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) up

# Откатить последнюю миграцию
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) down

# Показать статус миграций
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) status

# Создать новую миграцию
migrate-create:
	@read -p "Введите название миграции: " name; \
	goose -dir $(MIGRATIONS_DIR) create $${name} sql

# Полный сброс (откат всех миграций) - опасно!
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) reset

# Подключиться к БД через psql
db-connect:
	docker exec -it database psql -U $(APP_DB_USER) -d $(APP_DB_NAME)
run-producer:
	go run cmd/producer/main.go
run-app:
	go run cmd/app/main.go