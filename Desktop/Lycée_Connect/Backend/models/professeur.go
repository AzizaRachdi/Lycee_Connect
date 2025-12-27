package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Professeur struct {
	ID         string    `gorm:"type:char(36);primaryKey" json:"id"`
	UserID     string    `gorm:"type:char(36);not null" json:"user_id"`
	Specialite string    `gorm:"size:100;not null" json:"specialite"`
	Niveau     string    `gorm:"size:100;not null" json:"niveau"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *Professeur) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}
