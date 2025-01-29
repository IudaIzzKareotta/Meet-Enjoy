package service

import (
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	repository2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/repository"
)

type ParticipantsService struct {
	repos repository2.Participants
}

func (ps ParticipantsService) GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error) {
	return ps.repos.GetEventParticipants(eventId)
}

func (ps ParticipantsService) DeleteParticipant(userId int, eventId int, participantId int) error {
	return ps.repos.DeleteParticipant(userId, eventId, participantId)
}

func NewParticipantsService(repos repository2.Participants) *ParticipantsService {
	return &ParticipantsService{repos: repos}
}
