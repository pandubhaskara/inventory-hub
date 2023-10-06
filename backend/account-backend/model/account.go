package model

import (
	"time"
)

type Account struct {
	ID                     uint          `json:"id" gorm:"primary_key"`
	Email                  string        `json:"email" gorm:"unique"`
	Name                   string        `json:"name"`
	ProfilePicture         *string       `json:"profile_picture"`
	MobileNumber           *string       `json:"mobile_number"`
	Gender                 *string       `json:"gender"`
	MaritalStatus          *string       `json:"marital_status"`
	PlaceOfBirth           *string       `json:"place_of_birth"`
	DateOfBirth            *time.Time    `json:"date_of_birth"`
	NationalIdentityNumber *string       `json:"national_identity_number"`
	CreatedAt              time.Time     `json:"created_at"`
	CreatedBy              *string       `json:"created_by"`
	UpdatedAt              time.Time     `json:"updated_at"`
	UpdatedBy              *string       `json:"updated_by"`
	DeletedAt              *time.Time    `json:"deleted_at"`
	DeletedBy              *string       `json:"deleted_by"`
	Applications           []Application `json:"applications" gorm:"many2many:account_applications;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type AccountApplication struct {
	AccountID     uint `json:"account_id" gorm:"primary_key;autoIncrement:false"`
	ApplicationID uint `json:"application_id" gorm:"primary_key;autoIncrement:false"`
}
