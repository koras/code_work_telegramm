package services

import (
	"fmt"
	"math/rand"
	"moex/controllers"
	"moex/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
* Удаляем автора из исполнителей
 */
func deleteAutor(Performers []models.User, autor string) []string {
	// все авторы имеют собавку в начале
	autor = "@" + autor
	var newPerformers []string

	for _, performer := range Performers {
		fmt.Println(performer.Nickname)
		if autor != performer.Nickname {
			newPerformers = append(newPerformers, performer.Nickname)
		}
	}
	return newPerformers
}

/**
*
 */
func AssignTask(bot *tgbotapi.BotAPI, update tgbotapi.Update, stack string) {
	rand.Seed(time.Now().UnixNano())

	var Performer string
	var message string
	var preparePerformers []string
	var Performers []models.User
	// список авторов
	switch stack {
	case "go":
		Performers = controllers.GetPerformers("go")
	case "php":
		Performers = controllers.GetPerformers("php")
	}

	// автор сообщения
	autor := update.Message.From.UserName

	fmt.Println(preparePerformers)
	preparePerformers = deleteAutor(Performers, autor)

	fmt.Println(preparePerformers)
	if len(preparePerformers) > 0 {
		Performer = preparePerformers[rand.Intn(len(preparePerformers))]
		message = "@" + Performer + "\n Вам назначена задача на проверку качества кода #kk #" + stack
	} else {
		message = "Нет доступных исполнителей для проверки, надо указать стек технологий, php или go "
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

/**
* Запрашиваем исполнителей
 */
func GetPerformers(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	var message string
	var performersPHP []string
	var performersGO []string
	// список авторов
	Performers := controllers.AllImplementer()

	for _, performer := range Performers {
		fmt.Println(performer.Nickname + " " + performer.Lang)
		if performer.Lang == "php" {
			performersPHP = append(performersPHP, performer.Nickname)
		} else if performer.Lang == "go" {
			performersGO = append(performersGO, performer.Nickname)
		}
	}

	message = "Исполнители на проверку кода \n"

	if len(performersPHP) > 0 {
		message += "php: \n"
		for _, performerPHP := range performersPHP {
			fmt.Println(performerPHP)
			message += "@" + performerPHP + " \n"
		}
	}

	if len(performersGO) > 0 {
		message += "go: \n"
		for _, performerGO := range performersGO {
			fmt.Println(performerGO)
			message += "@" + performerGO + " \n"
		}
	}

	if len(Performers) == 0 {
		message = "Нет доступных исполнителей для проверки, надо указать стек технологий, php или go"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	bot.Send(msg)
}
