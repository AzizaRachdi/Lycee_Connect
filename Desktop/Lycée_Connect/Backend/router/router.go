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

	// CRUD professeurs
	apicrudprofs := r.Group("/lycee_Connect/backend")
	{
		prof := apicrudprofs.Group("/professeurs")
		{
			prof.POST("/", controllers.CreateProfesseur)
			prof.GET("/", controllers.GetProfesseurs)
			prof.GET("/:id", controllers.GetProfesseurByID)
			prof.PUT("/:id", controllers.UpdateProfesseur)
			prof.DELETE("/:id", controllers.DeleteProfesseur)
		}
	}
	// CRUD parents
	apicrudparents := r.Group("/lycee_Connect/backend")
	{
		parent := apicrudparents.Group("/parents")
		{
			parent.POST("/", controllers.CreateParent)
			parent.GET("/", controllers.GetParents)
			parent.GET("/:id", controllers.GetParentByID)
			parent.PUT("/:id", controllers.UpdateParent)
			parent.DELETE("/:id", controllers.DeleteParent)
		}
	}
	// CRUD élèves
	apicrudeleves := r.Group("/lycee_Connect/backend")
	{
		eleve := apicrudeleves.Group("/eleves")
		{
			eleve.POST("/", controllers.CreateEleve)
			eleve.GET("/", controllers.GetEleves)
			eleve.GET("/:id", controllers.GetEleveByID)
			eleve.PUT("/:id", controllers.UpdateEleve)
			eleve.DELETE("/:id", controllers.DeleteEleve)
		} 
	}

	return r
}
