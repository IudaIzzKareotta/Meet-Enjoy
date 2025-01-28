package repository

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	UsersTable        = os.Getenv("USERS_TABLE")
	EventsTable       = os.Getenv("EVENTS_TABLE")
	ParticipantsTable = os.Getenv("PARTICIPANTS_TABLE")
)

type Authorization interface {
	CreateUser(user MeetEnjoy2.User) (int, error)
	GetUser(username, password string) (MeetEnjoy2.User, error)
}

type Event interface {
	CreateEvent(event MeetEnjoy2.Event, userId int) (int, error)
	GetUserEvents(userId int) ([]MeetEnjoy2.Event, error)
	GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error)
	GetEventById(eventId int) (MeetEnjoy2.Event, error)
	UpdateEvent(eventId int, updateInput MeetEnjoy2.UpdateEventInput) error
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
