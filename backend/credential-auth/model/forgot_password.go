package model

import (
	"time"
)

type ForgotPassword struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	Token     string     `json:"token"`
	IssuedAt  time.Time  `json:"issued_at"`
	ExpiredAt time.Time  `json:"expired_at"`
	UsedAt    *time.Time `json:"used_at"`
	IsActive  bool       `json:"is_active"`
}
