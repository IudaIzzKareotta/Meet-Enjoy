package service

import (
	"MeetEnjoy"
	"MeetEnjoy/pkg/repository"
)

type EventService struct {
	repos repository.Event
}

func (es EventService) GetAllEvents(userId int) ([]MeetEnjoy.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (es EventService) CreateEvent(event MeetEnjoy.Event, userId int) (int, error) {
	return es.repos.CreateEvent(event, userId)
}

func (es EventService) GetUserEvents(userId int) ([]MeetEnjoy.Event, error) {
	return es.repos.GetUserEvents(userId)
}
func (es EventService) GetEventParticipants(eventId int) ([]MeetEnjoy.Participants, error) {
	return es.repos.GetEventParticipants(eventId)
}
func (es EventService) GetEventById(eventId int) (MeetEnjoy.Event, error) {
	return es.repos.GetEventById(eventId)
}

func (es EventService) UpdateEvent(eventId int, updateInput MeetEnjoy.UpdateEventInput) error {
	return es.repos.UpdateEvent(eventId, updateInput)
}

func (es EventService) DeleteEvent(eventId int, userId int) error {
	return es.repos.DeleteEvent(eventId, userId)
}

func NewEventService(repos repository.Event) *EventService {
	return &EventService{repos: repos}
}
