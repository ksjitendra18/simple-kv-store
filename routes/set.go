package routes

import (
	"net/http"
	"time"

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

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := c.MustGet("store").(*Store)

	expirySeconds := time.Duration(requestData.Options.Ex) * time.Second

	// store.Set(requestData.Key, requestData.Value, requestData.Options.Ex)
	store.Set(requestData.Key, requestData.Value, expirySeconds)
	c.JSON(200, gin.H{
		"message": "Data set successfully",
	})
}
