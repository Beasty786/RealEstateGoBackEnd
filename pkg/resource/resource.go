package resource

import (

	"net/http"


	"restate_backend/pkg/service"
	"github.com/gin-gonic/gin"


)

type Resource interface{
	CreateTables(context *gin.Context)

	HealthResource
	PropertyResource
	CategoryResource
}

type resourceImpl struct {
	srv service.Service
}

func NewResource(srv service.Service) *resourceImpl{
	return &resourceImpl{
		srv: srv,
	}
}

func (r resourceImpl) CreateTables(context *gin.Context) {
	r.srv.CreateTables()

	context.JSON(http.StatusCreated, gin.H{
		"message": "tables created",
	})
}

