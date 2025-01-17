package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"primaryKey"`
	Email         string `gorm:"unique;not null" validate:"required,email"`
	Password      string `gorm:"not null" validate:"required,min=6"`
	Name          string `gorm:"not null"`
	EmailVerified bool   `gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return
}
