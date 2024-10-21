package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Start() {
	r := gin.New()
	r.GET("/health", health)
	r.POST("/withBodyAndHeader", withBodyAndHeader)
	r.Run(":8082")
}

func health(c *gin.Context) {
	c.Writer.Write([]byte("Hello world!"))
}

type RequestBody struct {
	Message string `json:"message"`
}

func withBodyAndHeader(c *gin.Context) {
	apiKey := c.GetHeader("x-api-key")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing header x-xpi-key"})
		return
	}

	if _, err := uuid.Parse(apiKey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid x-api-key"})
		return
	}

	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON sent"})
		return
	}

	c.Status(http.StatusOK)
}
