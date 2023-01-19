package config

import (
	"fmt"
	"log"
	"moex/models"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir + ".env")
	environmentPath := filepath.Join(dir + ".env")
	errorENV := godotenv.Load(environmentPath)

	if errorENV != nil {
		panic("Failed to load env file ConnectDB")
	}

	dbUser := os.Getenv("DB_USER_WORK")
	dbPass := os.Getenv("DB_PASS_WORK")
	dbHost := os.Getenv("DB_HOST_WORK")
	dbName := os.Getenv("DB_NAME_WORK")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database where base  ")
	}

	db.AutoMigrate(&models.User{}, &models.Stack{}, &models.Appointment{})
	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
