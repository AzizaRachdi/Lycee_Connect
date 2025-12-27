package services

import (
	"errors"
	"fmt"

	"github.com/Lycee_Connect/config"
	"github.com/Lycee_Connect/models"
	"github.com/Lycee_Connect/security"
)

type AuthService struct{}

var Auth = AuthService{}

// Register avec roleID (UUID du rôle)
func (AuthService) Register(name, email, username, password, roleID string) (*models.User, error) {
	// Déclare la variable role
	var role models.Role

	// Vérifie que le rôle existe
	if err := config.DB.First(&role, "id = ?", roleID).Error; err != nil {
		return nil, errors.New("role invalide: UUID introuvable")
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: hash,
		RoleID:   roleID, // on utilise l'ID du rôle trouvé
	}

	if err := config.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (AuthService) Login(login, password string) (string, *models.User, error) {
	var u models.User
	if err := config.DB.Preload("Role").Where("email = ? OR username = ?", login, login).First(&u).Error; err != nil {
		return "", nil, errors.New("utilisateur introuvable")
	}
	if !security.CheckPassword(u.Password, password) {
		return "", nil, errors.New("mot de passe invalide")
	}
	// Mettre l'ID utilisateur et la description de rôle dans le token
	tok, err := security.GenerateToken(fmt.Sprintf("%s", u.ID), u.Role.Description)
	return tok, &u, err
}
