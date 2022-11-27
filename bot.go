package main

import (
	"moex/config"
	"moex/services"

	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {

}

func main() {
	defer config.DisconnectDB(db)
	// Вызываем бота
	services.TelegramBot()
}
