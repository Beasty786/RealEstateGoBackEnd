package service

import "fmt"

type PropertyService interface{
	GetAllProperties() (Properties, error)
	GetPropertyById(propertyId int) (*Property, error)
	GetPropertyByCategoryId(categoryId int) (Properties, error)
}

func (s serviceImpl) GetAllProperties() (Properties, error){
	repoProperties, err := s.repo.GetAllProperties()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	serviceProperties := Properties{}
	for _, property := range repoProperties{
		serviceProperties = append(
			serviceProperties, 
			fromRepoToServiceProperty(property),
		)
	}

	return serviceProperties, nil
}

func (s serviceImpl) GetPropertyById(propertyId int) (*Property, error){
	repoProperty, err := s.repo.GetPropertyById(propertyId)
	if err != nil {
		return nil, err
	}
	servProperty := fromRepoToServiceProperty(*repoProperty)
	return &servProperty,nil
}

func (s serviceImpl) GetPropertyByCategoryId(categoryId int) (Properties, error){
	repoProperties, err := s.repo.GetPropertiesByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}
	
	serviceProperties := Properties{}
	for _, property := range repoProperties{
		serviceProperties = append(
			serviceProperties, 
			fromRepoToServiceProperty(property),
		)
	}
	return serviceProperties,nil
}
