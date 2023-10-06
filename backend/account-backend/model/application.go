package model

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Icon      string         `json:"icon"`
	Url       string         `json:"url"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy *string        `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy *string        `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy *string        `json:"deleted_by"`
}
