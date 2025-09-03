package repository

import log "github.com/sirupsen/logrus"

type PropertyRepo interface{
	SaveProperty(property Property) error 
	GetAllProperties() ([]Property, error)
	GetPropertyById(id int) (*Property, error)
	DeletePropertyById(property_id int) (int, error)
	GetPropertiesByCategoryId(categoryId int) ([]Property, error)
	
}

func (r repository) SaveProperty(property Property) error {
	cat, err := r.GetCategoryByName(property.Name)
	if err != nil {
		return err
	}

	query := `
				INSERT INTO 
					properties (name, description, category_id, price, address, thumbnail_url, rating, num_baths, num_beds, size, owner_id, agent_id) 
				VALUES 
					(	:name,
						:description, 
						:category_id,
						:price,
						:address,
						:thumbnail_url,
						:rating, 
						:num_baths,
						:num_beds,
						:size,
						:owner_id,
						:agent_id
					);
				`
	insertData := map[string]interface{}{
		"name": property.Name,
		"description": property.Description,
		"category_id": &cat.Id,
		"address": property.Address,
		"thumbnail_url": property.Image_url,
		"price":property.Price,
		"rating": property.Rating,
		"num_baths": property.Num_baths,
		"num_beds": property.Num_beds,
		"size": property.Size,
		"owner_id": property.Owner_id,
		"agent_id": property.Agent_id,
	}

	_, err = r.db.NamedQuery(query, insertData)
	return err
}

func (r repository) GetAllProperties() ([]Property, error) {

	query := `
			SELECT
				property.id, property.name, description, category.name as category,price,  address,
				thumbnail_url, rating, num_baths, num_beds, size, owner_id, agent_id
			FROM 
				properties property
			INNER JOIN
				categories category on category.id = property.category_id;
	`
	properties := []Property{}

	err := r.db.Select(&properties, query); if err != nil{
		return nil, err
	}

	return properties, nil
}

func (r repository) GetPropertyById(id int) (*Property, error) {
	query := `
		SELECT
			property.id, property.name, description, category.name as category,price,  address,
			thumbnail_url, rating, num_baths, num_beds, size, owner_id, agent_id
		FROM 
			properties property
			INNER JOIN
			categories category on category.id = property.category_id
			WHERE property.id = $1;
	`
	property := Property{}

	err := r.db.Get(&property, query, id); if err != nil{
		log.Error("error occurred: ", err)
		return nil, err
	}

	return &property, nil
}

func (r repository) GetPropertiesByCategoryId(categoryId int) ([]Property, error) {
	query := `
		SELECT
			property.id, property.name, description, category.name as category,price,  address,
			thumbnail_url, rating, num_baths, num_beds, size, owner_id, agent_id
		FROM 
			properties property
			INNER JOIN
			categories category on category.id = property.category_id
			WHERE category.id = $1;
	`
	properties := []Property{}

	err := r.db.Select(&properties, query, categoryId); if err != nil{
		log.Error("error occurred: ", err)
		return nil, err
	}

	return properties, nil
}

func (r repository) DeletePropertyById(property_id int) (int, error) {
	query := `DELETE FROM properties WHERE id = $1`

	affectRows, err := r.db.MustExec(query, property_id).RowsAffected()

	return int(affectRows) , err
}
