package controllers

import (
	"net/http"

	"github.com/Lycee_Connect/models"
	"github.com/Lycee_Connect/services"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Role.Create(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func GetRoles(c *gin.Context) {
	roles, err := services.Role.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := services.Role.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rôle introuvable"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var updated models.Role
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Role.Update(id, &updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rôle mis à jour"})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := services.Role.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "rôle supprimé"})
}
