package service

import "backend/internal/repository"

type MainRouteService interface {
	GetData() (map[string]string, error)
}

type Service struct {
	MainRouteService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		MainRouteService: NewMainRouteService(repo),
	}
}
