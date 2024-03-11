package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MainRouteRepo interface {
}

type Repository struct {
	MainRouteRepo
}

func NewRepository(db *gorm.DB, rb *redis.Client) *Repository {
	return &Repository{
		MainRouteRepo: NewMainRoute(db, rb),
	}
}
