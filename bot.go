package main

import (
	"fmt"
	"log"
	"moex/services"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var token string

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("work.env"); err != nil {
		log.Print("No work.env file found")
	} else {
		token = os.Getenv("TOKEN_TELEGRAM")
	}
}

// https://habr.com/ru/post/351060/

func telegramBot() {

	fmt.Println("check")
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
		fmt.Println("error 6")
		panic(err)
	}
	// //Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	// //Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		word_kkk, _ := regexp.Compile("([k]{3}|[к]{3})")
		stackRegx, _ := regexp.Compile("([go]{2}|[php]{3})")
		stack := stackRegx.FindString(update.Message.Text)

		if word_kkk.MatchString(update.Message.Text) {
			services.AssignTask(bot, update, stack)
		}

		if update.Message.Text == "/help" || update.Message.Text == "/help@StartKKWork1Bot" {
			services.Help(bot, update)
		}
	}
}

func main() {
	// Вызываем бота
	telegramBot()
}
