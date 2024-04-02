package service

import "backend/internal/repository"

type MainRouteService interface {
	GetData() (map[string]string, error)
	// UpdateData For test)
	UpdateData()
	DeleteData()
}

type Service struct {
	MainRouteService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		MainRouteService: NewMainRouteService(repo),
	}
}
