package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email    string
	Password string
	Fistname string
	Lastname string
	gorm.Model
}
