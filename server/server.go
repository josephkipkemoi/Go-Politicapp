package server

import (
	"github.com/gin-gonic/gin"
)

// Set up router and set necessary configurations, headers, api routes etc
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(SetHeaders())

	// Set up API routes
	Routers(r)

	return r
}
