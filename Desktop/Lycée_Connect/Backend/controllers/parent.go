package controllers

import (
	"net/http"

	"github.com/Lycee_Connect/models"
	"github.com/Lycee_Connect/services"
	"github.com/gin-gonic/gin"
)

func CreateParent(c *gin.Context) {
	var parent models.Parent
	if err := c.ShouldBindJSON(&parent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Parent.Create(&parent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, parent)
}

func GetParents(c *gin.Context) {
	parents, err := services.Parent.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parents)
}

func GetParentByID(c *gin.Context) {
	id := c.Param("id")
	parent, err := services.Parent.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "parent introuvable"})
		return
	}
	c.JSON(http.StatusOK, parent)
}

func UpdateParent(c *gin.Context) {
	id := c.Param("id")
	var updated models.Parent
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Parent.Update(id, &updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "parent mis à jour"})
}

func DeleteParent(c *gin.Context) {
	id := c.Param("id")
	if err := services.Parent.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "parent supprimé"})
}
