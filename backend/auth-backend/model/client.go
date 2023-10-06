package model

type Client struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
