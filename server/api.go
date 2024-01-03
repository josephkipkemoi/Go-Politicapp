package server

import (
	"github.com/josephkipkemoi/politicapp/controllers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	public := r.Group("api/v1")

	public.GET("/", controllers.LandingController)
	public.POST("/register", controllers.Register)
}
