package models

type Role struct {
	Id		uint	`gorm:"primary_key; increment" json:"id"`
	Name	string	`gorm:"size:255; not null" json:"name"`
}
