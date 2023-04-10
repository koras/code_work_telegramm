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
	//autor = "@" + autor
	var newPerformers []string

	fmt.Println("=======")
	for _, performer := range Performers {
		fmt.Println(performer.Nickname + " != " + autor)
		if autor != performer.Nickname {
			newPerformers = append(newPerformers, performer.Nickname)
		}
	}
	return newPerformers
}

func AssignTask(bot *tgbotapi.BotAPI, update tgbotapi.Update, lang string) {
	
	rand.Seed(time.Now().UnixNano())

	var Performer string
	var message string
	var preparePerformers []string

	stack := controllers.GetStackByLang(lang)
	if stack.ID == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверно указан стек.")
		bot.Send(msg)
		return
	}

	// список авторов
	Performers := controllers.GetActivePerformers(stack)

	// автор сообщения
	author := update.Message.From.UserName

	fmt.Println(preparePerformers)
	preparePerformers = deleteAutor(Performers, author)

	fmt.Println(" author" + author)
	fmt.Println(preparePerformers)
	if len(preparePerformers) > 0 {
		Performer = preparePerformers[rand.Intn(len(preparePerformers))]
		if(update.Message.Chat.ID == -1001400698397){
			controllers.AddAppointment(Performer)
			fmt.Println("Добавляем в статистику")
		}else{
			fmt.Println("Вызван из другой группы")
		}
		
		message = "@" + Performer + "\n Вам назначена задача на проверку качества кода #kk #" + lang
	} else {
		message = "Нет доступных исполнителей для проверки"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

/**
* Запрашиваем исполнителей
 */
func GetPerformers(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var message = "Исполнители на проверку кода \n"
	var stacks map[string][]string

	stacks = controllers.AllImplementer()

	for stack, performers := range stacks {
		message += fmt.Sprintf("%s: \n", stack)
		for _, user := range performers {
			message += "" + user + " \n"
		}
	}

	fmt.Println(update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	bot.Send(msg)
}
