package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" validate:"required,min=5,max=30"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100)" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:varchar(255)" json:"password" validate:"required,min=8,max=30"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	if !isPasswordHashed(u.Password) {
		h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		fmt.Println(u.Password)
		u.Password = string(h)
	}
	return nil
}

func isPasswordHashed(password string) bool {
	return len(password) > 0 && (password[:4] == "$2a$" || password[:4] == "$2b$" || password[:4] == "$2y$")
}
