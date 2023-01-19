package controllers

import (
	"moex/config"
	"moex/models"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

// получаем всех пользователей
func GetPerformers(stack *models.Stack) []models.User {
	var users []models.User
	//TODO: реализовать в синтаксисе GORM
	db.Raw("select * from users where id in (select user_id from user_stacks where stack_id = 22) and now() not between vacation_from and vacation_to", stack.ID).Scan(&users)
	return users
}

// получаем пользователей не в отпуске
func GetActivePerformers(stack *models.Stack) []models.User {
	var users []models.User
	//TODO: реализовать в синтаксисе GORM
	db.Raw("select * from users where id in (select user_id from user_stacks where stack_id = ?)", stack.ID).Scan(&users)
	return users
}

func AddStackIfNotExists(stackLang string) *models.Stack {
	stack := models.Stack{}
	db.Model(&models.Stack{}).Where("lang = ?", stackLang).Find(&stack)
	if stack.ID == 0 {
		db.Model(&models.Stack{}).Create(&models.Stack{Lang: stackLang})
	}
	db.Model(&models.Stack{}).Where("lang = ?", stackLang).Find(&stack)
	return &stack
}

func AddPerformer(nickname string) *models.User {
	newUser := models.User{Nickname: nickname}
	db.Model(&models.User{}).Create(&newUser)

	return &newUser
}

func EditPerformer(nickname string, stacks []string) *models.User {
	var user models.User

	db.Where("nickname = ?", nickname).Find(&user)
	db.Model(&user).Association("Stacks").Replace([]models.Stack{})
	for _, lang := range stacks {
		stack := AddStackIfNotExists(lang)
		db.Model(&user).Association("Stacks").Append(stack)
	}

	return &user
}

func SetVacation(user *models.User, start time.Time, end time.Time) {
	user.VacationFrom = &start
	user.VacationTo = &end
	db.Model(&user).Save(&user)
}

// получаем список исполнителей под каждый стек
func AllImplementer() map[string][]string {
	var stackList = make(map[string][]string)
	var stacks []models.Stack

	db.Model(&models.Stack{}).Find(&stacks)

	var performers []models.User
	db.Model(&models.User{}).Preload("Stacks").Find(&performers)

	for _, user := range performers {
		for _, stack := range stacks {
			for _, curStack := range user.Stacks {
				if stack.Lang == curStack.Lang {
					stackList[stack.Lang] = append(stackList[stack.Lang], user.Nickname)
				}
			}
		}
	}

	return stackList
}

func DelUser(nickname string) bool {
	var user models.User

	db.Model(&models.User{}).Where("nickname = ?", nickname).Find(&user)
	if user.ID == 0 {
		return false
	}
	db.Exec("delete from user_stacks where user_id = ?", user.ID)
	db.Model(&models.Appointment{}).Where("user_id = ?", user.ID).Delete(&models.Appointment{})
	db.Delete(&user)
	return true
}
