package controllers

import "moex/models"

func GetStackByLang(lang string) *models.Stack {
	var stack models.Stack

	db.Model(&stack).Where("lang = ?", lang).Find(&stack)
	return &stack
}
