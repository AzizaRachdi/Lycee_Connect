package controllers

import (
    "net/http"

    "github.com/Lycee_Connect/models"
    "github.com/Lycee_Connect/services"
    "github.com/gin-gonic/gin"
)

func CreateProfesseur(c *gin.Context) {
    var prof models.Professeur
    if err := c.ShouldBindJSON(&prof); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
        return
    }
    if err := services.Professeur.Create(&prof); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, prof)
}

func GetProfesseurs(c *gin.Context) {
    profs, err := services.Professeur.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, profs)
}

func GetProfesseurByID(c *gin.Context) {
    id := c.Param("id")
    prof, err := services.Professeur.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "professeur introuvable"})
        return
    }
    c.JSON(http.StatusOK, prof)
}

func UpdateProfesseur(c *gin.Context) {
    id := c.Param("id")
    var updated models.Professeur
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
        return
    }
    if err := services.Professeur.Update(id, &updated); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "professeur mis à jour"})
}

func DeleteProfesseur(c *gin.Context) {
    id := c.Param("id")
    if err := services.Professeur.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "professeur supprimé"})
}
