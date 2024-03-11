package repository

import "gorm.io/gorm"

type MainRouteRepo interface {
}

type Repository struct {
	MainRouteRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		MainRouteRepo: NewMainRoute(db),
	}
}
