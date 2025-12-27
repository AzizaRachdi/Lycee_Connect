package services

import (
    "github.com/Lycee_Connect/config"
    "github.com/Lycee_Connect/models"
)

type ProfesseurService struct{}

var Professeur = ProfesseurService{}

func (ProfesseurService) Create(p *models.Professeur) error {
    return config.DB.Create(p).Error
}

func (ProfesseurService) GetAll() ([]models.Professeur, error) {
    var profs []models.Professeur
    err := config.DB.Find(&profs).Error
    return profs, err
}

func (ProfesseurService) GetByID(id string) (*models.Professeur, error) {
    var prof models.Professeur
    err := config.DB.First(&prof, "id = ?", id).Error
    return &prof, err
}

func (ProfesseurService) Update(id string, updated *models.Professeur) error {
    return config.DB.Model(&models.Professeur{}).Where("id = ?", id).Updates(updated).Error
}

func (ProfesseurService) Delete(id string) error {
    return config.DB.Delete(&models.Professeur{}, "id = ?", id).Error
}
