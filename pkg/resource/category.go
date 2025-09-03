package resource

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryResource interface {
	GetAllCategories(c *gin.Context)
	
}

func (r resourceImpl) GetAllCategories(c *gin.Context){
	categories, err := r.srv.GetAllCategories()

	if err != nil {
		log.Error("error getting all categories: %", err)
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"message": fmt.Errorf("an error occured getting categories: %w", err),
			"data": nil,
		})	
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "fetched categories successfully",
		"data": categories,
	})
}