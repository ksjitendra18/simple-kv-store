package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ksjitendra18/simple-kv-store/routes"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	server := gin.Default()
	server.Use(corsMiddleware())
	store := routes.NewStore()

	// Add middleware to inject the store instance into each request context
	server.Use(func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	})

	routes.RegisterRoutes(server)

	server.Run("localhost:8080")
}
