package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "health ok"})
}
