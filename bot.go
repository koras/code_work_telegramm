package main

import (
	"fmt"
	"moex/config"
	"moex/services"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	//"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {

	fmt.Println("init")

	env := ".env"

	errorENV := godotenv.Load(env)
	//errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file init")
	}
	//	db = config.ConnectDB()
}

func main() {
	defer config.DisconnectDB(db)
	// Вызываем бота
	services.TelegramBot()
}
