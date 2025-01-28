package MeetEnjoy

type User struct {
	Id       int    `json:"-" db:"id"`                                 // Айдишник
	Username string `json:"username" db:"username" binding:"required"` // Ник
	Name     string `json:"name" db:"name" binding:"required"`         // Имя
	Surname  string `json:"surname" db:"surname" binding:"required"`   // Фамилия
	Email    string `json:"email" db:"email" binding:"required"`       // Почта
	Password string `json:"password" binding:"required"`               // Пароль
}
