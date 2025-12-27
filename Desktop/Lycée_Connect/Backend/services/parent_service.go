package services

import (
    "github.com/Lycee_Connect/config"
    "github.com/Lycee_Connect/models"
)

type ParentService struct{}

var Parent = ParentService{}

func (ParentService) Create(p *models.Parent) error {
    return config.DB.Create(p).Error
}

func (ParentService) GetAll() ([]models.Parent, error) {
    var parents []models.Parent
    err := config.DB.Find(&parents).Error
    return parents, err
}

func (ParentService) GetByID(id string) (*models.Parent, error) {
    var parent models.Parent
    err := config.DB.First(&parent, "id = ?", id).Error
    return &parent, err
}

func (ParentService) Update(id string, updated *models.Parent) error {
    return config.DB.Model(&models.Parent{}).Where("id = ?", id).Updates(updated).Error
}

func (ParentService) Delete(id string) error {
    return config.DB.Delete(&models.Parent{}, "id = ?", id).Error
}
