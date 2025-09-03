package resource

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type PropertyResource interface{
	GetAllProperties(c *gin.Context)
	GetPropertyById(c *gin.Context)
	GetPropertiesByCategoryId(c *gin.Context)
}

func (r resourceImpl) GetAllProperties(c *gin.Context){
	properties, err := r.srv.GetAllProperties()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
		gin.H{
			"status": http.StatusNotFound,
			"message": fmt.Errorf("%w", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": properties,
	})
}

func (r resourceImpl) GetPropertyById(c *gin.Context){
	propertyId, err := strconv.Atoi(c.Param("propertyIs"))
	if err != nil{
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status" : http.StatusBadRequest,
				"message": "parameter isn't and integer or simply not passed",
			}, 
		)
		return
	}

	property, err := r.srv.GetPropertyById(propertyId)

	if err != nil {
		log.Error("error from getting property by id: ", err)
		c.AbortWithStatusJSON(http.StatusNotFound,
		gin.H{
			"status": http.StatusNotFound,
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": property,
	})

}

func (r resourceImpl) GetPropertiesByCategoryId(c *gin.Context){
	strCategoryId := c.Param("categoryId")
	if strCategoryId == "" {
		err := fmt.Errorf("no category Id passed")
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"status": http.StatusBadRequest,
			"message": "couldn't get categoryId",
		})
		return
	}

	log.Info("Getting properties for category id: ", strCategoryId)
	categoryId, err := strconv.Atoi(strCategoryId)
	if err != nil{
		log.Error(err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status" : http.StatusBadRequest,
				"message": "parameter isn't and integer or simply not passed",
			}, 
		)
		return
	}
	
	properties, err := r.srv.GetPropertyByCategoryId(categoryId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
		gin.H{
			"status": http.StatusNotFound,
			"message": fmt.Errorf("%w", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": properties,
	})	
}
