package controllers

import (
	"fmt"
	"moex/config"
	"moex/models"

	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

// получаем пользователей для назначения задания
func GetPerformers(lang string) []models.User {
	var users []models.User
	if err := db.Joins("INNER JOIN stacks ON stacks.users_id = users.id and stacks.lang=? ", lang).Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return users
}

// получаем список исполнителей
func AllImplementer() []models.Appointments {
	// SELECT  * FROM `stacks` JOIN `users`  ON `users`.id = `stacks`.users_id ORDER BY `lang`, `nickname`
	var appointments []models.Appointments
	if err := db.Select("nickname", "stacks.lang").Table("users").Joins("INNER JOIN stacks ON stacks.users_id = users.id").Order("lang desc, nickname").Find(&appointments).Error; err != nil {
		fmt.Println(err)
	}
	return appointments
}

// добавляем информацию для статистики, сколько и кому что было ранее назначено
func SetAppointment(appointment models.Appointment) {
	db.Select("Lang", "UsersID").Create(appointment)
}

func GetImplementer() []models.User {
	var users []models.User
	if err := db.Joins("INNER JOIN stacks ON stacks.users_id = users.id").Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return users
}
