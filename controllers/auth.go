package controllers

import (
	"net/http"

	"github.com/Lycee_Connect/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var in struct{ Name, Email, Username, Password, Role string }
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	u, err := services.Auth.Register(in.Name, in.Email, in.Username, in.Password, in.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}
func Login(c *gin.Context) {
	var in struct{ Login, Password string }
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalide"})
		return
	}
	tok, u, err := services.Auth.Login(in.Login, in.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tok, "role": u.Role})
}
