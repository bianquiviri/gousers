package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Name            string         `gorm:"size:255;not null;index" json:"name"`
	Email           string         `gorm:"size:255;unique;not null" json:"email"`
	EmailVerifiedAt *time.Time      `json:"email_verified_at"`
	Password        string         `gorm:"size:255;not null" json:"-"`
	Address         string         `gorm:"size:255" json:"address"`
	PhoneNumber     string         `gorm:"size:255" json:"phone_number"`
	DateOfBirth     *time.Time      `json:"date_of_birth"`
	Gender          string         `gorm:"size:255" json:"gender"`
	RememberToken   string         `gorm:"size:100" json:"-"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
