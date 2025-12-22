package router

import (
	"github.com/Lycee_Connect/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)
	return r
}
