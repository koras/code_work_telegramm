package services

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Исполнители на проверку качества кода PHP
var PerformersPHP = []string{
	"@foxxy_9",
	"@deSatEm",
	"@Hezurus",
	"@avp365",
	"@mikhail",
	"@avp365",
	"@ivanstukalin",
	"@liigos",
}

// Исполнители на проверку качества кода GO
var PerformersGo = []string{
	"@deSatEm",
	"@avp365",
	"@mikhail",
	"@avp365",
	"@ivanstukalin",
	"@akaTheEnd",
	"@liigos",
}

/**
* Удаляем автора из исполнителей
 */
func deleteAutor(Performers []string, autor string) []string {
	// все авторы имеют собавку в начале
	autor = "@" + autor
	var newPerformers []string
	for _, performer := range Performers {
		if autor != performer {
			newPerformers = append(newPerformers, performer)

		}
	}
	return newPerformers
}

/**
*  Помощь
 */
func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	welcome := "Привет\n"
	welcome += "Я бот который ставит задачу на проверку качества кода\n"
	welcome += "Для назначения проверки достаточно упомянуть в чате 'ккк' или 'kkk'\n"
	welcome += "Так же 'не обязательно' можно упомянуть язык 'php' или 'go'\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcome)
	bot.Send(msg)
}

/**
* Парсер за каждй день
 */
func AssignTask(bot *tgbotapi.BotAPI, update tgbotapi.Update, stack string) {
	var Performer string
	// список авторов
	Performers := PerformersPHP
	// автор сообщения
	autor := update.Message.From.UserName

	switch stack {
	case "go":
		Performers = PerformersGo
	case "php":
		Performers = PerformersPHP
	}
	// удаляем автора задачи
	Performers = deleteAutor(Performers, autor)

	Performer = Performers[rand.Intn(len(Performers))]

	message := Performer + "\n Вам назначена задача на проверку качества кода " + stack
	//	+		"\n Автор задачи " + autor

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

// https://ru.stackoverflow.com/questions/527779/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B8%D1%82%D1%8C-%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82-%D0%B8%D0%B7-%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2%D0%B0-%D0%BF%D0%BE-%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D1%8E

/*
/help@StartKKWork1Bot - Помощь, кто я и за чем
/performers@StartKKWork1Bot - Исполнители на задачу

*/
