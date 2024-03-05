package repository

import "gorm.io/gorm"

type mainRoutePostgres struct {
	db *gorm.DB
}

func NewMainRoute(db *gorm.DB) *mainRoutePostgres {
	return &mainRoutePostgres{db: db}
}
