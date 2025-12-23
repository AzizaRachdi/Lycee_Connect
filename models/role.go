package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          string `gorm:"type:char(36);primaryKey" json:"id"`
	Code        int    `gorm:"uniqueIndex" json:"code"`
	Description string `gorm:"size:50;not null" json:"description"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}
