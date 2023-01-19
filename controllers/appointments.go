package controllers

import (
	"moex/models"
	"time"
)

func AddAppointment(nickname string) {
	var user models.User
	db.Model(&models.User{}).Where("nickname = ?", nickname).Find(&user)
	datetime := time.Now()
	appointment := models.Appointment{
		CreatedAt: &datetime,
		UserID:    user.ID,
	}
	db.Create(&appointment)
}

func GetStats() map[string]int64 {
	var users []models.User
	stats := make(map[string]int64)

	db.Model(models.User{}).Find(&users)
	for _, user := range users {
		var count int64
		db.Model(&models.Appointment{}).Where("user_id = ?", user.ID).Count(&count)

		stats[user.Nickname] = count
	}
	return stats
}
