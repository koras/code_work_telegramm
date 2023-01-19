package models

import "time"

type User struct {
	ID           uint
	Nickname     string
	Stacks       []Stack       `gorm:"many2many:user_stacks;constraint:OnDelete:CASCADE"`
	Appointments []Appointment `gorm:"foreignKey:UserID;"`
	VacationFrom *time.Time
	VacationTo   *time.Time
}

type Stack struct {
	ID   uint `gorm:"primarykey"`
	Lang string
}

// Для вычисления кому и что назначено
type Appointment struct {
	ID        uint
	CreatedAt *time.Time
	UserID    uint
}

// Схема исполнителя для запроса их списка
type Performers struct {
	StackID uint
	Users   []string
}

func (User) TableName() string {
	return "users"
}

func (Stack) TableName() string {
	return "stacks"
}

func (Appointment) TableName() string {
	return "appointments"
}
