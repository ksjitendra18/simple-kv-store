package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SetRequest struct {
	Key     string   `json:"key"`
	Value   string   `json:"value"`
	Options *Options `json:"options,omitempty"`
}

type Options struct {
	Ex int `json:"ex,omitempty"`
}

func Set(c *gin.Context) {

	var requestData SetRequest

	// Bind the JSON data to the SetRequest struct
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("key", requestData.Key, requestData.Value)

	// Get the store instance
	store := c.MustGet("store").(*Store)
	store.Set(requestData.Key, requestData.Value)
	c.JSON(200, gin.H{
		"message": "Data set successfully",
	})
}
