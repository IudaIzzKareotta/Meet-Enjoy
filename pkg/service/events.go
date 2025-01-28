package service

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	repository2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/repository"
)

type EventService struct {
	repos repository2.Event
}

func (es EventService) CreateEvent(event MeetEnjoy2.Event, userId int) (int, error) {
	return es.repos.CreateEvent(event, userId)
}

func (es EventService) GetUserEvents(userId int) ([]MeetEnjoy2.Event, error) {
	return es.repos.GetUserEvents(userId)
}
func (es EventService) GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error) {
	return es.repos.GetEventParticipants(eventId)
}
func (es EventService) GetEventById(eventId int) (MeetEnjoy2.Event, error) {
	return es.repos.GetEventById(eventId)
}

func (es EventService) UpdateEvent(eventId int, updateInput MeetEnjoy2.UpdateEventInput) error {
	return es.repos.UpdateEvent(eventId, updateInput)
}

func (es EventService) DeleteEvent(eventId int, userId int) error {
	return es.repos.DeleteEvent(eventId, userId)
}

func NewEventService(repos repository2.Event) *EventService {
	return &EventService{repos: repos}
}
