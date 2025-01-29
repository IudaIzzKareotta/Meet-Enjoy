package service

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	repository2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/repository"
)

type Authorization interface {
	CreateUser(user MeetEnjoy2.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

type Events interface {
	CreateEvent(event MeetEnjoy2.Event, userId int) (int, error)
	GetUserEvents(userId int) ([]MeetEnjoy2.Event, error)
	GetEventById(eventId int) (MeetEnjoy2.Event, error)
	UpdateEvent(eventId int, updateInput MeetEnjoy2.UpdateEventInput) error
	DeleteEvent(eventId int, userId int) error
}

type Participants interface {
	GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error)
	DeleteParticipant(userId int, eventId int, participantId int) error
}

type Service struct {
	Authorization
	Events
	Participants
}

func NewService(repos *repository2.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Events:        NewEventService(repos),
		Participants:  NewParticipantsService(repos),
	}
}
