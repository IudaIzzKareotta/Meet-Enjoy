package repository

import (
	"MeetEnjoy"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (ap *AuthPostgres) CreateUser(user MeetEnjoy.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (username, name, surname, email, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id")

	row := ap.db.QueryRow(query, user.Username, user.Name, user.Surname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		log.Printf("Error executing query: %s, error: %v", query, err)
		return -1, err
	}

	return id, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (MeetEnjoy.User, error) {
	var user MeetEnjoy.User
	query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")
	err := ap.db.Get(&user, query, username, password)

	return user, err
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
