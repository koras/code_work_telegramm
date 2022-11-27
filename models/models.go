package models

// список сотрудников

type User struct {
	ID       uint
	Nickname string
	Stacks   Stack `json:"Stack" gorm:"foreignKey:UsersID"`
}

// Какому сотруднику и какой стек
type Stack struct {
	ID      uint
	UsersID uint
	Lang    string
}

// Для вычисления кому и что назначено
type Appointment struct {
	ID      uint
	UsersID uint
	Lang    string
}

// Исполнители по списку
type Appointments struct {
	UsersID  uint
	Lang     string
	Nickname string
}

func (User) TableName() string {
	return "users"
}

func (Stack) TableName() string {
	return "stacks"
}
