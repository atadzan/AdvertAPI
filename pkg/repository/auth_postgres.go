package repository

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func(r *AuthPostgres) CreateUser(user AdvertAPI.SignUpInput)(int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, phone_number, created_at) VALUES($1, $2, $3, $4) RETURNING id",
		usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.PhoneNumber, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func(r *AuthPostgres) GetUser(username, password string)(AdvertAPI.User, error){
	var user AdvertAPI.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
