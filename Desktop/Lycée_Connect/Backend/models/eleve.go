package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Eleve struct {
	ID            string    `gorm:"type:char(36);primaryKey" json:"id"`
	UserID        string    `gorm:"type:char(36);not null" json:"user_id"`
	ParentID      string    `gorm:"type:char(36);not null" json:"parent_id"`
	Classe        string    `gorm:"size:50;not null" json:"classe"`
	DateNaissance time.Time `json:"date_naissance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (e *Eleve) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
