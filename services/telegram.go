package services

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string

// https://habr.com/ru/post/351060/

func TelegramBot() {

	errorENV := godotenv.Load(".env")

	if errorENV != nil {
		panic("No .env file found")
	} else {
		token = os.Getenv("TOKEN_TELEGRAM")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
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
			AssignTask(bot, update, stack)
		}

		if update.Message.Text == "/performers" || update.Message.Text == "/performers@StartKKWork1Bot" {
			GetPerformers(bot, update)
		}
		if update.Message.Text == "/help" || update.Message.Text == "/help@StartKKWork1Bot" {
			Help(bot, update)
		}
	}
}
