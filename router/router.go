package router

import (
	"github.com/Lycee_Connect/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/lycee_Connect/backend")
	auth.POST("/auth/register", controllers.Register)
	auth.POST("/auth/login", controllers.Login)

	// Nouveau groupe pour le CRUD des rôles
	apicrudroles := r.Group("/lycee_Connect/backend")
	{
		role := apicrudroles.Group("/roles")
		{
			role.POST("/", controllers.CreateRole)
			role.GET("/", controllers.GetRoles)
			role.GET("/:id", controllers.GetRoleByID)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:id", controllers.DeleteRole)
		}
	}
	return r
}
