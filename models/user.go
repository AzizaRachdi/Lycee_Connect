package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




// Modèle User
type User struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Password  string    `json:"-"` // mot de passe hashé, jamais renvoyé en clair
	RoleID    string    `gorm:"type:char(36);not null" json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Hook pour générer un UUID avant insertion
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}


