package server

import (
	"io"

	"github.com/gin-gonic/gin"
)

func (r routes) saveToken(rg *gin.RouterGroup) {
	token := rg.Group("/saveToken")

	token.GET("/", saveFunction)
	token.PUT("/", saveFunction1)
}

func saveFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "no token",
	})
}

func saveFunction1(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": string(body),
	})
}
