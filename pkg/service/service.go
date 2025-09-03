package service

import (
	"restate_backend/pkg/repository"
)

type Service interface{
	Health() error
	CreateTables()

	PropertyService
	CategoryService
}

type serviceImpl struct {
	repo repository.Repository
}


func NewService(repo repository.Repository) *serviceImpl{
	return &serviceImpl{
		repo: repo,
	}
}

func (s serviceImpl) CreateTables() {
	s.repo.DeleteTables()
	s.repo.CreateTables()
}

