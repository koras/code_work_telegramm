package services

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string

// https://habr.com/ru/post/351060/

func TelegramBot() {

	//	errorENV := godotenv.Load()

	//if errorENV != nil {
	//	panic("No .env file found")
	//} else {
	token = os.Getenv("TOKEN_TELEGRAM_WORK")
	//	}

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
		stackRegex, _ := regexp.Compile("[^ ,][^ ,]*")

		if word_kkk.MatchString(update.Message.Text) {
			stack_start := word_kkk.FindStringIndex(update.Message.Text)[1] + 1
			var stack string
			if stack_start <= len(update.Message.Text)-1 {
				stack = stackRegex.FindString(update.Message.Text[stack_start:])
			} else {
				stack = ""
			}
			fmt.Printf("\"%s\"", stack)
			AssignTask(bot, update, stack)
		}

		if update.Message.Text == "/performers" || update.Message.Text == "/performers@StartKKWork1Bot" {
			GetPerformers(bot, update)
		}

		if update.Message.Text == "/help" || update.Message.Text == "/help@StartKKWork1Bot" {
			Help(bot, update)
		}

		if len(update.Message.Text) > 4 && update.Message.Text[:4] == "/set" {
			var nickname string
			var stack = []string{}

			nickname = strings.Split(update.Message.Text[5:], " ")[0]
			if nickname[0] == '@' {
				nickname = nickname[1:]
			}

			if len(strings.Split(update.Message.Text[5:], " ")) > 1 {
				stack = strings.Split(strings.Split(update.Message.Text[5:], " ")[1], ",")
			}

			EditOrAddUser(bot, update, nickname, stack)
		}

		if len(update.Message.Text) > 9 && update.Message.Text[:9] == "/vacation" {
			var nickname string
			var start string
			var duration int

			nickname = strings.Split(update.Message.Text[10:], " ")[0]
			if nickname[0] == '@' {
				nickname = nickname[1:]
			}

			if len(strings.Split(update.Message.Text[10:], " ")) == 3 {
				start = strings.Split(update.Message.Text[10:], " ")[1]
				duration, _ = strconv.Atoi(strings.Split(update.Message.Text[10:], " ")[2])
			} else {
				duration = 0
			}

			ScheduleVacation(bot, update, nickname, start, duration)
		}

		if len(update.Message.Text) > 5 && update.Message.Text[:5] == "/kick" {
			var nickname string

			nickname = strings.Split(update.Message.Text[6:], " ")[0]
			if nickname[0] == '@' {
				nickname = nickname[1:]
			}

			DeleteUser(bot, update, nickname)
		}

		if update.Message.Text == "/stats" || update.Message.Text == "/performers@StartKKWork1Bot" {
			GetStatistics(bot, update)
		}
	}
}
