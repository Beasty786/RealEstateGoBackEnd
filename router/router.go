package router

import (
	"restate_backend/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
 	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/create-tables", init.Resource.CreateTables)

		health := api.Group("/health")
		health.GET("", init.Resource.HealthCheck)	

		property := api.Group("/properties")
		property.GET("/", init.Resource.GetAllProperties)
		property.GET("/:propertyId", init.Resource.GetPropertyById)
		property.GET("/by-category/:categoryId", init.Resource.GetPropertiesByCategoryId)

		category := api.Group("/categories")
		category.GET("", init.Resource.GetAllCategories)
	}
	
	return router
}
