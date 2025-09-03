package model

type Property struct {
	Name int `json:"name"`
	Owner_id int `json:"owner_id"`
	Agent_id int `json:"agent_id"`
	Description string `json:"description"`
	Category int `json:"category"`
	Address string `json:"address"`
	Image_url string `json:"thumbnail_url"`
	Rating float32 `json:"rating"`
	Num_baths int `json:"num_baths"`
	Num_beds int `json:"num_beds"`
	Size int `json:"size"`
}

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
