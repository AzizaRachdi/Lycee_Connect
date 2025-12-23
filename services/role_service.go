package services

import (
    "github.com/Lycee_Connect/config"
    "github.com/Lycee_Connect/models"
)

type RoleService struct{}

var Role = RoleService{}

func (RoleService) Create(role *models.Role) error {
    return config.DB.Create(role).Error
}

func (RoleService) GetAll() ([]models.Role, error) {
    var roles []models.Role
    err := config.DB.Find(&roles).Error
    return roles, err
}

func (RoleService) GetByID(id string) (*models.Role, error) {
    var role models.Role
    err := config.DB.First(&role, "id = ?", id).Error
    return &role, err
}

func (RoleService) Update(id string, updated *models.Role) error {
    return config.DB.Model(&models.Role{}).Where("id = ?", id).Updates(updated).Error
}

func (RoleService) Delete(id string) error {
    return config.DB.Delete(&models.Role{}, "id = ?", id).Error
}
