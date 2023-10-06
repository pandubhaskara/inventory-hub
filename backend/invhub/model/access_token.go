package model

import (
	"time"
)

type AccessToken struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	ClientID  int        `json:"client_id"`
	IssuedAt  time.Time  `json:"issued_at"`
	ExpiredAt time.Time  `json:"expired_at"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
