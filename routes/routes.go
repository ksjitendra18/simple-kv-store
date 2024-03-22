package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	v0Routes := server.Group("/v0")

	v0Routes.GET("/healthcheck", HealthCheck)
	v0Routes.POST("/set", Set)
	v0Routes.GET("/get/:key", Get)

}
