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

func (AuthService) Register(name, email, username, password, role string) (*models.User, error) {
	hash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &models.User{Name: name, Email: email, Username: username, Password: hash, Role: models.Role(role)}
	if err := config.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func (AuthService) Login(login, password string) (string, *models.User, error) {
	var u models.User
	if err := config.DB.Where("email = ? OR username = ?", login, login).First(&u).Error; err != nil {
		return "", nil, errors.New("utilisateur introuvable")
	}
	if !security.CheckPassword(u.Password, password) {
		return "", nil, errors.New("mot de passe invalide")
	}
	tok, err := security.GenerateToken(fmt.Sprintf("%d", u.ID), string(u.Role))
	return tok, &u, err
}
