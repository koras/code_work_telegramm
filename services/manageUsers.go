package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"moex/config"
	"moex/controllers"
	"moex/models"
	"strings"
	"time"
)

var db *gorm.DB = config.ConnectDB()
var months = map[string]string{"01": "Jan", "02": "Feb", "03": "Mar", "04": "Apr", "05": "May", "06": "Jun", "07": "Jul", "08": "Aug", "09": "Sep", "10": "Oct", "11": "Nov", "12": "Dec"}

func EditOrAddUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, nickname string, stacks []string) {
	var user *models.User
	var message string

	db.Model(&models.User{}).Where("nickname = ?", nickname).Find(&user)
	if user.ID == 0 {
		message = fmt.Sprintf("Пользователь @%s успешно добавлен", nickname)
		user = controllers.AddPerformer(nickname)
	}
	controllers.EditPerformer(nickname, stacks)
	if message == "" {
		message = fmt.Sprintf("Пользователь @%s успешно изменен", user.Nickname)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

func ScheduleVacation(bot *tgbotapi.BotAPI, update tgbotapi.Update, nickname string, start string, duration int) {
	var user *models.User
	var message string
	var startTimestamp time.Time
	var endTimestamp time.Time

	db.Model(&models.User{}).Where("nickname = ?", nickname).Find(&user)
	if user.ID == 0 {
		message = fmt.Sprintf("Пользователь @%s не найден", nickname)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		bot.Send(msg)
		return
	}
	pieces := strings.Split(start, ".")
	startString := fmt.Sprintf("%s-%s-%s", pieces[2], months[pieces[1]], pieces[0])
	format := "2006-Jan-02"
	startTimestamp, err := time.Parse(format, startString)
	if err != nil {
		message = fmt.Sprintf("Неверная дата %s", nickname)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		bot.Send(msg)
		return
	}
	endTimestamp = startTimestamp.AddDate(0, 0, duration)
	fmt.Println(startTimestamp, endTimestamp)
	controllers.SetVacation(user, startTimestamp, endTimestamp)

	if message == "" {
		message = fmt.Sprintf("Отпуск для пользователя @%s установлен с %d.%d.%d по %d.%d.%d", user.Nickname, startTimestamp.Day(), startTimestamp.Month(), startTimestamp.Year(), endTimestamp.Day(), endTimestamp.Month(), endTimestamp.Year())
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}

func DeleteUser(bot *tgbotapi.BotAPI, update tgbotapi.Update, nickname string) {
	var message string
	success := controllers.DelUser(nickname)
	if !success {
		message = "Пользоваетля не существует"
	} else {
		message = fmt.Sprintf("Пользователь @%s удален", nickname)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	bot.Send(msg)
}
