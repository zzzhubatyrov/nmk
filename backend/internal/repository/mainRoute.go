package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type mainRoutePostgres struct {
	db *gorm.DB
	rb *redis.Client
}

func NewMainRoute(db *gorm.DB, rb *redis.Client) *mainRoutePostgres {
	return &mainRoutePostgres{db: db, rb: rb}
}
