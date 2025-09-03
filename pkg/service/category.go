package service

import (
	"fmt"

)

type CategoryService interface {
	GetAllCategories() (Categories, error)
}

func (s serviceImpl) GetAllCategories() (Categories, error) {
	serviceCategories := Categories{}

	repoCategories, err := s.repo.GetAllCategories()

	if err != nil {
		return nil, fmt.Errorf("error in service layer %w", err)
	}

	for _, category := range repoCategories{
		serviceCategories = append(serviceCategories, fromrepoToServiceCategory(category))
	}

	return serviceCategories, nil
}