package model

import "time"

type Product struct {
	ID          uint       `json:"id"`
	InventoryID int        `json:"inventory_id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Price       int        `json:"price"`
	Type        string     `json:"type"`
	Quantity    int        `json:"quantity"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *string    `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	UpdatedBy   *string    `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	DeletedBy   *string    `json:"-"`
}
