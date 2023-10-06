package model

import "time"

type Inventory struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Location  *string    `json:"location"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	UpdatedBy *string    `json:"-"`
	DeletedAt *time.Time `json:"-"`
	DeletedBy *string    `json:"-"`
}
