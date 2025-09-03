package service

import (
	"restate_backend/pkg/repository"
)

type Property struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
	Address string `json:"address"`
	Image_url string `json:"thumbnail_url"`
	Rating float32 `json:"rating"`
	Num_baths float32 `json:"num_baths"`
	Num_beds int `json:"num_beds"`
	Size int `json:"size"`
	Price float32 `json:"price"`
	Owner_id int `json:"owner_id"`
	Agent_id int `json:"agent_id"`
}

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Categories []Category

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image_url string `json:"image_url"`
}

type Properties []Property

func fromRepoToServiceProperty(repoProperty repository.Property) Property {
	return Property{
		Id: repoProperty.Id,
		Name: repoProperty.Name,
		Description: repoProperty.Description,
		Category: repoProperty.Category,
		Address: repoProperty.Address,
		Image_url: repoProperty.Image_url,
		Rating: repoProperty.Rating,
		Num_baths: repoProperty.Num_baths,
		Num_beds: repoProperty.Num_beds,
		Size: repoProperty.Size,
		Owner_id: repoProperty.Owner_id,
		Agent_id: repoProperty.Agent_id,
	}
}

func fromrepoToServiceCategory(repoCategory repository.Category) Category {
	return Category{
		Id: repoCategory.Id,
		Name: repoCategory.Name,
	}
}