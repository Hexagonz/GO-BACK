package models

import (
	"time"
)

type RefreshToken struct {
	SessionID     uint      `gorm:"primaryKey" json:"session_id"`
	Refresh_Token string    `gorm:"type:text" json:"refresh_token" validate:"required"`
	UserAgent     string    `gorm:"type:varchar(255)" json:"user_agent" validate:"required"`
	ExpiredAt	  time.Time	`validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
