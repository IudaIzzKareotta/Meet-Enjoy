package repository

import (
	"fmt"
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type EventPostgres struct {
	db *sqlx.DB
}

func (ep EventPostgres) CreateEvent(event MeetEnjoy2.Event, userId int) (int, error) {
	var eventId int

	tx, err := ep.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
    INSERT INTO events (title, content, author_id, photo_url, event_date, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, CURRENT_DATE, CURRENT_DATE) RETURNING id`
	row := tx.QueryRow(query, event.Title, event.Content, userId, event.PhotoUrl, event.EventDate)
	if err := row.Scan(&eventId); err != nil {
		return 0, err
	}

	query = `INSERT INTO participants (user_id, event_id, current_status, status_updated_at) VALUES ($1, $2, $3, CURRENT_DATE)`
	_, err = tx.Exec(query, userId, eventId, "Yes")
	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (ep EventPostgres) GetUserEvents(userId int) ([]MeetEnjoy2.Event, error) {
	var events []MeetEnjoy2.Event

	query := "SELECT title, content, photo_url, event_date, created_at, updated_at FROM events WHERE author_id = $1"
	err := ep.db.Select(&events, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return events, nil
}

func (ep EventPostgres) GetEventParticipants(eventId int) ([]MeetEnjoy2.Participants, error) {
	var participants []MeetEnjoy2.Participants
	query := "SELECT user_id, event_id, current_status, status_updated_at FROM participants WHERE event_id = $1"

	err := ep.db.Select(&participants, query, eventId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return participants, nil
}

func (ep EventPostgres) GetEventById(eventId int) (MeetEnjoy2.Event, error) {
	var event MeetEnjoy2.Event
	query := "SELECT * FROM events WHERE id = $1"
	err := ep.db.Get(&event, query, eventId)
	if err != nil {
		return event, fmt.Errorf("failed to execute query: %w", err)
	}
	return event, nil
}

func (ep EventPostgres) UpdateEvent(eventId int, updateInput MeetEnjoy2.UpdateEventInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateInput.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, updateInput.Title)
		argId++
	}

	if updateInput.Content != "" {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, updateInput.Content)
		argId++
	}

	if updateInput.PhotoUrl != "" {
		setValues = append(setValues, fmt.Sprintf("photo_url=$%d", argId))
		args = append(args, updateInput.PhotoUrl)
		argId++
	}

	if updateInput.EventDate != "" {
		eventDate, err := time.Parse(time.RFC3339, updateInput.EventDate)
		if err != nil {
			return fmt.Errorf("invalid event date format: %v", err)
		}
		setValues = append(setValues, fmt.Sprintf("event_date=$%d", argId))
		args = append(args, eventDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE events SET %s, updated_at = NOW() WHERE id = $%d`, setQuery, argId)

	args = append(args, eventId)

	_, err := ep.db.Exec(query, args...)
	return err
}

func (ep EventPostgres) DeleteEvent(eventId int, userId int) error {
	query := `DELETE FROM events WHERE id = $1 AND author_id = $2`
	_, err := ep.db.Exec(query, eventId, userId)
	if err != nil {
		return err
	}

	return err
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}
