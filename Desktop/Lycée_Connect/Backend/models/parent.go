package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Parent struct {
    ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
    UserID    string    `gorm:"type:char(36);not null" json:"user_id"`
    Phone     string    `gorm:"size:20;not null" json:"phone"`
    Address   string    `gorm:"type:text;not null" json:"address"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (p *Parent) BeforeCreate(tx *gorm.DB) (err error) {
    p.ID = uuid.New().String()
    return
}
