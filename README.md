## 📦 Демонстрационный сервис заказов

---

Сервис обрабатывает заказы из Kafka и предоставляет API для получения данных о заказах с кэшированием в Redis.

### Поток данных:
0. При запуске все данные из БД кэшируются в Redis
1. Заказы поступают из Kafka и сохраняются в PostgreSQL
2. При запросе пользователя сначала проверяется кэш Redis
3. Если данных нет в кэше - берутся из БД и кэшируются
4. Ответ возвращается пользователю

### Инструкция:
0. Создать .env файл и заполнить поля, как в .env.example файле. Можно просто скопировать данные из .env.example и вставить в .env файл
1. Поднять контейнеры с помощью `docker-compose run`.
2. Запустить основное приложение `make run-app`
3. Запустить скрипт, который будет посылать заказы в сервис, с помощью `make run-producer`
4. Открыть `web/index.html`. Для тестирования доступны следующие id товаров:
    - `b563feb7b2b84b6test`
    - `a1b2c3d4e5f67890test`
    - `f0e1d2c3b4a59687test`
    - `test9876543210zyx`
    - `a1b2c3d4e5f67890rand`
    - `b6c7d8e9f0a1b2c3rand`
    - `c9d8e7f6a5b4c3d2rand`
    - `d3e4f5a6b7c8d9e0rand`
5. swagger можно запустить по `http://localhost:8081/swagger`

### Архитектура проекта:
- `cmd/app` - стартовая точка самого приложения, `cmd/producer` - скрипт, который отправляет заранее определённые заказы в Kafka. Сами заказы определены в `cmd/producer/testModels`
- `db/init` - скрипт создания роли в базе данных
- `docs` - здесь swagger хранится
- `internal/app` - здесь хранятся инициализация соединений Redis, Kafka, PostgreSql, сам сервер, который запросы к API принимает, logger и реализация gracefull shutdown
- `internal/broker/consumer/kafka` - логика consumer'a
- `internal/configs` - получение данных из .env файла
- `internal/controllers/handler` - обработчик запросов. `internal/controllers/response` - ответы на запрос
-  `internal/domain` - модель заказа и её валидация
- `internal/lib/logger` - библиотека для красивого вывода логов из `slog/log`. Код позаимствовал из проект `https://github.com/GolangLessons/url-shortener`
- `internal/repository/cache` - работа с Redis. `internal/repository/converter` - конвертирует модель бизнес логики в модель репозитория (бд). `internal/repository/db` - работа с базой данных. `internal/repository/model` - модель репозитория (бд)
- `internal/service` - сама бизнес логика, связанная с кэшированием, загрузкой и получением заказа.
- `web` - хранится index.html.


### Использованные библиотеки
- Работа с Kafka: `github.com/confluentinc/confluent-kafka-go`
- Работа с Redis: `github.com/redis/go-redis/v9`
- Работа с Postgres: `gorm.io/gorm` и `gorm.io/driver/postgres`
- Веб-сервер: `github.com/go-chi/chi/v5`