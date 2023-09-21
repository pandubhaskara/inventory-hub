package model

import (
	"time"
)

type User struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Email      string     `json:"email" gorm:"unique"`
	Password   string     `json:"-"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *string    `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	UpdatedBy  *string    `json:"-"`
	DeletedAt  *time.Time `json:"-"`
	DeletedBy  *string    `json:"-"`
}
