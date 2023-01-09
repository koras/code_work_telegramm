package main

import (
	"moex/config"
	"moex/services"

	_ "github.com/joho/godotenv/autoload"

	//"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	// Вызываем бота
	services.TelegramBot()
}
