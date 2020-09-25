package models

import (
	"errors"
	"example/config"
	"fmt"
	"github.com/badoux/checkmail"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {

	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Take(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	err = config.DB.Where("id =?", id).Take(&User{}).UpdateColumn(
		map[string]interface{}{
			"name":      user.Name,
			"phone":     user.Phone,
			"email":     user.Email,
			"update_at": time.Now(),
			"address":   user.Address,
		},
	).Take(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func GetUserByUsernameAndPassword(user *User, username string, password string) (err error) {
	if err = config.DB.Where("username = ? ", username).Find(user).Error; err != nil {
		return err
	}
    fmt.Println(username, password, user.Password)
	if err = VerifyPassword(user.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Listen when create
func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	fmt.Println("da vao hash", user)
	return nil
}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if user.Name == "" {
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		if user.Password == "" {
			return errors.New("required password")
		}
		if user.Username == "" {
			return errors.New("required username")
		}
		return nil
	case "register":
		if user.Name == "" {
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		if user.Password == "" {
			return errors.New("required password")
		}
		if user.Username == "" {
			return errors.New("required username")
		}
		return nil
	case "update":
		if user.Name == "" {
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if user.Password == "" {
			return errors.New("required password")
		}
		if user.Username == "" {
			return errors.New("required username")
		}
		return nil
	default:
		if user.Name == "" {
			return errors.New("required name")
		}
		if user.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

func (user *User) Prepare() {
	user.Id = 0
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
}
