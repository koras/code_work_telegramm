package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"moex/controllers"
)

func GetStatistics(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var message string
	stats := make(map[string]int64)
	stats = controllers.GetStats()

	message = "Статистика по количеству проверок:\n"
	for user, count := range stats {
		message += fmt.Sprintf("@%s: %d проверок\n", user, count)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
