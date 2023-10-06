package model

import (
	"time"
)

type RefreshToken struct {
	ID            uint       `json:"id" gorm:"primary_key"`
	AccessTokenID int        `json:"access_token_id"`
	IssuedAt      time.Time  `json:"issued_at"`
	ExpiredAt     time.Time  `json:"expired_at"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
