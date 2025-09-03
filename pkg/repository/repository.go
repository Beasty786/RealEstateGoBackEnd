package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	HealthCheck() error
	CreateTables()
	DeleteTables()

	CategoryRepo
	PropertyRepo
	UserRepo
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func ConnectToDB(database string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=shameel password=#123^bak3nd dbname=%s sslmode=disable", database))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func (r repository) CreateTables() {
	query := `
		CREATE TABLE categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL -- Increased length for category names
		);

		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			image_url VARCHAR(255)
		);

		CREATE TABLE properties (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL, -- Increased length for property name/title
			description TEXT,
			category_id INTEGER NOT NULL,
			owner_id INTEGER NOT NULL,
			agent_id INTEGER NULL,
			price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
			address VARCHAR(255),
			thumbnail_url VARCHAR(255),
			rating DECIMAL(2,1),    -- Adjusted for 1-5 scale with half-point precision (e.g., 4.5)
			num_baths DECIMAL(2,1),
			num_beds INTEGER,
			size INTEGER,
			CONSTRAINT fk_owner
				FOREIGN KEY (owner_id)
				REFERENCES users (id)
				ON DELETE CASCADE,
			CONSTRAINT fk_agent
				FOREIGN KEY (agent_id)
				REFERENCES users (id)
				ON DELETE SET NULL,
			CONSTRAINT fk_category
				FOREIGN KEY (category_id)
				REFERENCES categories (id)
				ON DELETE RESTRICT -- Good: prevents deleting categories if properties exist
		);

		CREATE TABLE amenities (
			id SERIAL PRIMARY KEY,
			name VARCHAR(25) NOT NULL
		);

		CREATE TABLE property_amenity (
			property_id INTEGER NOT NULL,
			amenity_id INTEGER NOT NULL,
			PRIMARY KEY (property_id, amenity_id), -- Added composite primary key
			CONSTRAINT fk_property
				FOREIGN KEY (property_id)
				REFERENCES properties (id)
				ON DELETE CASCADE, -- If property deleted, remove its amenity associations
			CONSTRAINT fk_amenity
				FOREIGN KEY (amenity_id)
				REFERENCES amenities (id)
				ON DELETE RESTRICT -- Prevents deleting an amenity if properties still use it
		);

		CREATE TABLE property_images (
			id SERIAL PRIMARY KEY,
			image_url VARCHAR(255) NOT NULL, -- Image URL should always be present
			property_id INTEGER NOT NULL,   -- Image should always belong to a property
			CONSTRAINT fk_property
				FOREIGN KEY (property_id)
				REFERENCES properties (id) -- Corrected typo: REFERENCE -> REFERENCES
				ON DELETE CASCADE         -- If property deleted, delete its images
		);

		CREATE TABLE reviews (
			id SERIAL PRIMARY KEY,
			review TEXT NOT NULL,         -- Changed to TEXT for potentially longer reviews
			rating DECIMAL(2,1),          -- Adjusted for 1-5 scale with half-point precision
			property_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			CONSTRAINT fk_property        -- Corrected typo: CONSTRAIN -> CONSTRAINT
				FOREIGN KEY (property_id)
				REFERENCES properties(id)
				ON DELETE CASCADE,        -- If property deleted, delete its reviews
			CONSTRAINT fk_user            -- Corrected typo: CONSTRAINS -> CONSTRAINT
				FOREIGN KEY (user_id)
				REFERENCES users (id)
				ON DELETE CASCADE         -- If user deleted, delete their reviews
		);
	`

	r.db.MustExec(query)
}

func (r repository) DeleteTables() {
	query := `
	DROP TABLE IF EXISTS reviews;
	DROP TABLE IF EXISTS property_images;
	DROP TABLE IF EXISTS property_amenity;
	DROP TABLE IF EXISTS amenities;
	DROP TABLE IF EXISTS properties;
	DROP TABLE IF EXISTS categories;
	DROP TABLE IF EXISTS users; -- Drop users last if reviews cascade from it.
	`
	r.db.MustExec(query)
}
