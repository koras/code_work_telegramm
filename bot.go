package main

import (
	"moex/config"
	"moex/services"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}
}

func main() {
	defer config.DisconnectDB(db)
	// Вызываем бота
	services.TelegramBot()
}
