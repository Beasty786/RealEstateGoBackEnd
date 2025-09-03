package repository

import "fmt"

type CategoryRepo interface {
	SaveCategories(categories []string) error
	GetAllCategories() ([]Category, error)
	GetCategoryById(category_id int) (*Category, error)
	GetCategoryByName(name string) (*Category, error)
}

func (r repository) SaveCategories(categories []string) error {

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transcation inserting categories with error: %w", err)
	}

	defer func ()  {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				err = fmt.Errorf("failed to commit transcation: %w", err)
			}
		}
	}()

	query := `INSERT INTO categories (name) VALUES ($1);`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}


	for _, categoryName := range categories {
		_, err := stmt.Exec(categoryName)
		if err != nil {
			return fmt.Errorf("failed to insert: %s, with error: %w", categoryName, err)
		}
	}
	return nil
}

func (r repository) GetAllCategories() ([]Category, error) {
	query := `
			SELECT
				id, name
			FROM 
				categories;
				`
	categories := []Category{}

	err := r.db.Select(&categories, query); if err != nil {
		return nil, fmt.Errorf("failed to get all categories with error: %w",err)
	}

	return categories, nil
}

func (r repository) GetCategoryById(category_id int) (*Category, error) {
	query := `
			SELECT 
				id, name 
			FROM 
				categories
			WHERE 
				id = $1;
	`
	category := Category{}
	err := r.db.Get(&category, query, category_id); if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r repository) GetCategoryByName(name string) (*Category, error) {
	query := `
			SELECT 
				id, name 
			FROM 
				categories
			WHERE 
				id = 1;
	`
	category := Category{}
	err := r.db.Get(&category, query); if err != nil {
		return nil, err
	}
	return &category, nil
}