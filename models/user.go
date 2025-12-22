package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleProf      Role = "PROF"
	RoleEleve     Role = "ELEVE"
	RoleParent    Role = "PARENT"
	RoleDirecteur Role = "DIRECTEUR"
)

type User struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Avant l’insertion, générer automatiquement un UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
