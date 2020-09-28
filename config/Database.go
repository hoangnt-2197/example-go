package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

type DBConfig struct {
	Host string
	Port int
	DBName string
	User string
	Password string
}

func BuildDBConfig() *DBConfig {
	 dbConfig := DBConfig {
	 	 Host : GetEnv("Host", ""),
	 	 Port : GetEnvAsInt("Port", 3306),
	 	 User: GetEnv("User", ""),
	 	 DBName: GetEnv("DBName", ""),
	 	 Password: GetEnv("Password" , ""),
	 }

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func GetEnv(key string, defaultVal string ) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}