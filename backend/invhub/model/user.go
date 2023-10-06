package model

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	// Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"-"`
	IsAdmin   bool       `json:"is_admin" gorm:"defult:false"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	UpdatedBy *string    `json:"-"`
	DeletedAt *time.Time `json:"-"`
	DeletedBy *string    `json:"-"`
}
