package controllers

import (
	"github.com/gin-gonic/gin"
)

func LandingController(c *gin.Context) {
	message := "Go 1.21"

	c.JSON(200, gin.H{
		"message": message,
	})
}
