package app

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	db  *gorm.DB
	rdb *redis.Client
	ctx *context.Context
}

func New() (*Application, error) {
	db, err := newDatabase()
	if err != nil {
		return &Application{}, err
	}

	rdb, ctx, err := newRedis()
	if err != nil {
		return &Application{}, err
	}

	return &Application{
		db:  db,
		rdb: rdb,
		ctx: ctx,
	}, nil
}
