package main

import (
	"example/config"
	"example/models"
	"example/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"github.com/jinzhu/gorm"
)

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %s", err)
	}
}

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Role{}, &models.User{})
	r := routes.SetupRouter()
	r.Run()
}