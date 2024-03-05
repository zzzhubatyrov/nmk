package service

import "backend/internal/repository"

type mainRouteService struct {
	repo repository.MainRouteRepo
}

func NewMainRouteService(repo repository.MainRouteRepo) *mainRouteService {
	return &mainRouteService{repo: repo}
}

func (m *mainRouteService) GetData() (map[string]string, error) {
	return map[string]string{"Hello": "World!"}, nil
}
