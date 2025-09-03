package resource

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResource interface {
	HealthCheck(context *gin.Context) 
}


func (r resourceImpl) HealthCheck(context *gin.Context) {
	err := r.srv.Health()

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}