package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"type:string;not null" json:"role"` // Corrected the GORM tag for the Role field
}
