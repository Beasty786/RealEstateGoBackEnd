package config

import (
	"restate_backend/pkg/repository"
	"restate_backend/pkg/resource"
	"restate_backend/pkg/service"
)

type Initialization struct {
	Resource resource.Resource
}

func NewInitialization(repo repository.Repository,
	srv service.Service,
	resource resource.Resource,
) *Initialization {
	return &Initialization{
		Resource: resource,
	}
}

func Init() *Initialization {

	db := repository.ConnectToDB("restate")
	repositoryImpl := repository.NewRepository(db)
	serviceImpl := service.NewService(repositoryImpl)
	resourceImpl := resource.NewResource(serviceImpl)

	return &Initialization{
		Resource: *resourceImpl,
	}
}
