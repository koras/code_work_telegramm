package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
*  Помощь
 */
func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	welcome := "Привет\n"
	welcome += "Я бот который ставит задачу на проверку качества кода\n"
	welcome += "Для назначения проверки достаточно упомянуть в чате 'ккк' или 'kkk'\n"
	welcome += "Так же обязательно нужно упомянуть стек через пробел после kkk\n"
	welcome += "/set [nickname] [стек через запятую] создаст пользоавтеля с заданным стеком. При существовании пользователя обновит его стек\n"
	welcome += "/vacation [nickname] [dd.mm.yyyy] [days] запланирует пользователю отпуск с заданной даты на заданное количество дней\n"
	welcome += "/kick [nickname] удалит пользователя\n"
	welcome += "/stats сформирует статистику по выполненным проверкам\n"
	welcome += "/performers сформирует список исполнителей по технологиям\n"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcome)
	bot.Send(msg)
}

// https://ru.stackoverflow.com/questions/527779/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B8%D1%82%D1%8C-%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82-%D0%B8%D0%B7-%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2%D0%B0-%D0%BF%D0%BE-%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%B8%D1%8E

/*
/help@StartKKWork1Bot - Общая информация
/performers@StartKKWork1Bot - Исполнители на задачу

*/
