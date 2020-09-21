package models

import (
	"errors"
	"example/config"
	"fmt"
	"github.com/badoux/checkmail"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	 "strings"
)

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err !=nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
   if err = config.DB.Where("id = ?", id).First(user).Error; err != nil{
	   return err
   } 
   return nil
}

func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error){
	fmt.Println(user)
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func GetUserByUsernameAndPassword(user *User) (err error) {
	if err = config.DB.Where("username = ? AND password = ?", user.Username , user.Password).Find(user).Error; err !=nil {
		return err
	}
	return nil
}

func Hash(password string) ([]byte, error) {
  return  bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword,password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error{
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Validate(action string) error  {
	switch strings.ToLower(action) {
	case "update":
		if user.Name== ""{
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
        if err := checkmail.ValidateFormat(user.Email); err!=nil {
        	return errors.New("invalid email")
		}
		return nil
	case "login":
		if user.Password== ""{
			return errors.New("required password")
		}
		if user.Username == "" {
			return errors.New("required username")
		}
		return nil
	default:
		if user.Name== ""{
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err!=nil {
			return errors.New("invalid email")
		}
		return nil
	}

}


