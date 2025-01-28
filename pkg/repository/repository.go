package repository

import (
	"MeetEnjoy"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	UsersTable        = os.Getenv("USERS_TABLE")
	EventsTable       = os.Getenv("EVENTS_TABLE")
	ParticipantsTable = os.Getenv("PARTICIPANTS_TABLE")
)

type Authorization interface {
	CreateUser(user MeetEnjoy.User) (int, error)
	GetUser(username, password string) (MeetEnjoy.User, error)
}

type Event interface {
	CreateEvent(event MeetEnjoy.Event, userId int) (int, error)
	GetUserEvents(userId int) ([]MeetEnjoy.Event, error)
	GetEventParticipants(eventId int) ([]MeetEnjoy.Participants, error)
	GetEventById(eventId int) (MeetEnjoy.Event, error)
	UpdateEvent(eventId int, updateInput MeetEnjoy.UpdateEventInput) error
	DeleteEvent(eventId int, userId int) error
}

type Repository struct {
	Authorization
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Event:         NewEventPostgres(db),
	}
}
