package models

type User struct {
	Id			uint 	`json:"id"`
	Name		string 	`json:"name"`
	Email		string	`json:"email"`
	Phone		string	`json:"phone"`
	Address		string	`json:"address"` 
	Username 	string 	`json:"username"`
	Password 	string	`json:"password"`
}

func (b *User)  TableName() string{
	return "user"
}
