package controllers

import (
	"net/http"

	"github.com/Lycee_Connect/models"
	"github.com/Lycee_Connect/services"
	"github.com/gin-gonic/gin"
)

func CreateEleve(c *gin.Context) {
	var eleve models.Eleve
	if err := c.ShouldBindJSON(&eleve); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Eleve.Create(&eleve); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, eleve)
}

func GetEleves(c *gin.Context) {
	eleves, err := services.Eleve.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eleves)
}

func GetEleveByID(c *gin.Context) {
	id := c.Param("id")
	eleve, err := services.Eleve.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "élève introuvable"})
		return
	}
	c.JSON(http.StatusOK, eleve)
}

func UpdateEleve(c *gin.Context) {
	id := c.Param("id")
	var updated models.Eleve
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	if err := services.Eleve.Update(id, &updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "élève mis à jour"})
}

func DeleteEleve(c *gin.Context) {
	id := c.Param("id")
	if err := services.Eleve.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "élève supprimé"})
}
