package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"restate_backend/pkg/repository"
)

var _ = Describe("Property", func() {
	db := repository.ConnectToDB("restate_app")
	repo := repository.NewRepository(db)

	BeforeEach(func() {
		repo.DeleteTables()
		repo.CreateTables()
		db.Exec(`INSERT INTO users (name) VALUES ('tests_user');`)
		repo.SaveCategories([]string{"House", "Condo"})
		properties := []repository.Property{
				{
					Name: "Best House",
					Description: "",
					Category: "House",
					Price: 8900,
					Address: "1234 House Street",
					Image_url: "image_url",
					Rating: 4.5,
					Num_baths: 3,
					Num_beds: 3,
					Size: 400,
					Owner_id: 1,
					Agent_id: 1,
				},
				{
					Name: "Best Condo",
					Description: "",
					Category: "Condo",
					Address: "1234 Condo Street",
					Image_url: "image_url",
					Price: 2400,
					Rating: 4.2,
					Num_baths: 3,
					Num_beds: 3,
					Size: 400,
					Owner_id: 1,
					Agent_id: 1,
				},
			}
		for _, property := range properties {
			_ = repo.SaveProperty(property)
		} 
	})

	Context("SaveProperty: Inserting property data into database", func ()  {
		It("should insert multiple recodrs at a given time", func ()  {
			properties := []repository.Property{
				{
					Name: "Best Apartment",
					Description: "",
					Category: "Apartment",
					Price:8300,
					Address: "1234 Apartment Street",
					Image_url: "image_url_1",
					Rating: 4.5,
					Num_baths: 3,
					Num_beds: 3,
					Size: 400,
					Owner_id: 1,
					Agent_id: 1,
				},
				{
					Name: "Best Condo",
					Description: "",
					Category: "Condo",
					Address: "1234 Condo Street",
					Image_url: "image_url",
					Price: 2400,
					Rating: 4.2,
					Num_baths: 3,
					Num_beds: 3,
					Size: 400,
					Owner_id: 1,
					Agent_id: 1,
				},
			}
			for _, property := range properties {
				saveErr := repo.SaveProperty(property)
				Expect(saveErr).To(BeNil())

			} 
			returnRepo, err := repo.GetAllProperties()
			Expect(err).To(BeNil())
			Expect(returnRepo).To(HaveLen(len(properties)+2))
		})
	})

	Context("GetAllProperties(): should get all properties", func ()  {
		It("should get all properties from the database", func() {
			properties, err := repo.GetAllProperties()
			Expect(err).To(BeNil())
			Expect(properties).To(Not(BeEmpty()))
		})
	})
	
	Context("GetPropertyById(): get a single property", func ()  {
		It("should get property by id", func ()  {
			propertyId := 1
			property, err := repo.GetPropertyById(propertyId)

			Expect(err).To(BeNil())
			Expect(property).To(Not(BeNil()))
			Expect(property.Id).To(Equal(propertyId))
		})
	})

	Context("DeletePropertyById(): delete a single property by its unique id", func ()  {
		It("should get a delete", func ()  {
			propertyId := 1
			expectedAffectedRow := 1
			rowAffected, err := repo.DeletePropertyById(propertyId)
			Expect(rowAffected).To(Equal(expectedAffectedRow))
			Expect(err).To(BeNil())
		})

		It("should return zero affected rows", func() {
			propertyId := 5
			expectedAffectedRow := 0
			rowAffected, err := repo.DeletePropertyById(propertyId)
			Expect(rowAffected).To(Equal(expectedAffectedRow))
			Expect(err).To(BeNil())
		})
	})
})
