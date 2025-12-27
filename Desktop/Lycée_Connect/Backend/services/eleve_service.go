package services

import (
	"github.com/Lycee_Connect/config"
	"github.com/Lycee_Connect/models"
)

type EleveService struct{}

var Eleve = EleveService{}

func (EleveService) Create(e *models.Eleve) error {
	return config.DB.Create(e).Error
}

func (EleveService) GetAll() ([]models.Eleve, error) {
	var eleves []models.Eleve
	err := config.DB.Find(&eleves).Error
	return eleves, err
}

func (EleveService) GetByID(id string) (*models.Eleve, error) {
	var eleve models.Eleve
	err := config.DB.First(&eleve, "id = ?", id).Error
	return &eleve, err
}

func (EleveService) Update(id string, updated *models.Eleve) error {
	return config.DB.Model(&models.Eleve{}).Where("id = ?", id).Updates(updated).Error
}

func (EleveService) Delete(id string) error {
	return config.DB.Delete(&models.Eleve{}, "id = ?", id).Error
}
