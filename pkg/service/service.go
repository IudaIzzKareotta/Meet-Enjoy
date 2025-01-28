package service

import (
	"MeetEnjoy"
	"MeetEnjoy/pkg/repository"
)

type Authorization interface {
	CreateUser(user MeetEnjoy.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

type Events interface {
	CreateEvent(event MeetEnjoy.Event, userId int) (int, error)
	GetUserEvents(userId int) ([]MeetEnjoy.Event, error)
	GetEventParticipants(eventId int) ([]MeetEnjoy.Participants, error)
	GetEventById(eventId int) (MeetEnjoy.Event, error)
	UpdateEvent(eventId int, updateInput MeetEnjoy.UpdateEventInput) error
	DeleteEvent(eventId int, userId int) error
}

type Service struct {
	Authorization
	Events
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Events:        NewEventService(repos),
	}
}
