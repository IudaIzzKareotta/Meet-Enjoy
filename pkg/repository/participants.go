package repository

import (
	"fmt"
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type ParticipantsPostgres struct {
	db *sqlx.DB
}

func (p ParticipantsPostgres) saveInvite(invite MeetEnjoy2.Invite) error {
	query := "INSERT INTO invites (event_id, token, created_at, expires_at) VALUES ($1, $2, $3, $4)"
	_, err := p.db.Exec(query, invite.Id, invite.Token, time.Now(), time.Hour*12)
	if err != nil {
		logrus.Errorf("Error insert invite %s:", err)
		return err
	}

	return nil
}

func (p ParticipantsPostgres) GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error) {
	var participants []MeetEnjoy2.Participants
	query := "SELECT user_id, event_id, current_status, status_updated_at FROM participants WHERE event_id = $1"

	err := p.db.Select(&participants, query, eventId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return participants, nil
}

func (p ParticipantsPostgres) DeleteParticipant(userId int, eventId int, participantId int) error {
	query := "DELETE FROM participants p USING events e WHERE p.event_id = $1 AND e.author_id = $2 AND p.user_id = $3"
	_, err := p.db.Exec(query, eventId, userId, participantId)
	if err != nil {
		logrus.Errorf("failed to execute query: %s", err)
		return err
	}

	return err
}

func NewParticipantsPostgres(db *sqlx.DB) *ParticipantsPostgres {
	return &ParticipantsPostgres{db: db}
}
