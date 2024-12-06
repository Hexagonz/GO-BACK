package models

import (
	"crypto/sha256"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"username" validate:"required,min=5,max=30"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100)" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:varchar(255)" json:"password" validate:"required,min=8,max=30"`
	CreatedAt time.Time
	UpdatedAt time.Time 
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	h := sha256.Sum256([]byte(u.Email))
	u.Password = fmt.Sprintf("%x", h[:])
	return nil
}

func (u *Users) BeforeSave(tx *gorm.DB) (err error) {
	h := sha256.Sum256([]byte(u.Email))
	u.Password = fmt.Sprintf("%x", h[:])
	return nil
}
