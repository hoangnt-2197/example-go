package models

import "time"

type User struct {
	Id			uint 		`gorm:"primary_key; auto_increment" json:"id"`
	Name		string 		`gorm:"size:255; not null" json:"name"`
	Email		string		`gorm:"size:100; not null; unique" json:"email"`
	Phone		string		`json:"phone"`
	Address		string		`gorm:"size:255" json:"address"`
	Username 	string 		`gorm:"size:100; not null; unique" json:"username"`
	Password 	string		`gorm:"size:100;not null;" json:"password"`
	CreatedAt 	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

type UserResponse struct {
	Id			uint 	`json:"id"`
	Name		string 	`json:"name"`
	Email		string	`json:"email"`
	Phone		string	`json:"phone"`
	Address		string	`json:"address"`
	Username 	string 	`json:"username"`
}

func (user *User)  TableName() string{
	return "user"
}
