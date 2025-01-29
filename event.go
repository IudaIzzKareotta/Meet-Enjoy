package MeetEnjoy

import "time"

type Participants struct {
	UserId          int       `json:"user_id" db:"user_id" binding:"required"`       // Айдишник юзера
	EventId         int       `json:"event_id" db:"event_id" binding:"required"`     // Айди ивента
	Status          string    `json:"status" db:"current_status"`                    // Статус пойдет или нет
	CreatedAt       time.Time `json:"created_id" db:"created_at" binding:"required"` // Дата вступления в ивент
	StatusUpdatedAt time.Time `json:"status_updated_at" db:"status_updated_at"`      // Дата обновления статуса юзера
}

type Event struct {
	Id        int       `json:"-" db:"id"`                                     // Айдишник
	Title     string    `json:"title" db:"title" binding:"required"`           // Айдишник создателя
	AuthorId  int       `json:"author_id" db:"author_id"`                      // Название
	Content   string    `json:"content" db:"content" binding:"required"`       // Текст
	PhotoUrl  string    `json:"photo_url" db:"photo_url"`                      // Фотка
	EventDate time.Time `json:"event_date" db:"event_date" binding:"required"` // Дата мероприятия
	CreatedAt time.Time `json:"created_at" db:"created_at"`                    // Дата создания
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`                    // Дата обновления
}

type UpdateEventInput struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PhotoUrl  string `json:"photo_url"`
	EventDate string `json:"event_date"`
}

type Invite struct {
	Id        uint      `json:"id"`
	EventId   int       `json:"event_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
