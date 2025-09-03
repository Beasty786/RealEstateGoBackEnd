package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"restate_backend/pkg/repository"
)

var _ = Describe("Category", func() {
	db := repository.ConnectToDB("restate_app")
	repo := repository.NewRepository(db)

	BeforeEach(func ()  {
		repo.DeleteTables()
		repo.CreateTables()
	})

	Context("SaveCategories", func ()  {
		It("should save categories", func ()  {
			categories := []string{"Home", "Shack"}
			err := repo.SaveCategories(categories)
			Expect(err).To(BeNil())
			insertedCategories, err := repo.GetAllCategories()
			Expect(err).To(BeNil())
			Expect(insertedCategories).To(HaveLen(len(categories)))
		})
	})
})
