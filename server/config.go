package server

import "github.com/gin-gonic/gin"

func SetHeaders() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json:charset=utf-8")
		c.Header("Host", c.Request.Host)
		c.Header("X-Powered-By", "go/1.21")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
}
