package repository

type Property struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Description string `db:"description"`
	Category string `db:"category"`
	Address string `db:"address"`
	Image_url string `db:"thumbnail_url"`
	Rating float32 `db:"rating"`
	Num_baths float32 `db:"num_baths"`
	Num_beds int `db:"num_beds"`
	Size int `db:"size"`
	Price float32 `db:"price"`
	Owner_id int `db:"owner_id"`
	Agent_id int `db:"agent_id"`
}

type Category struct {
	Id int `db:"id"`
	Name string `db:"name"`
}

type User struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Image_url string `db:"image_url"`
}