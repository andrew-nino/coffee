package repository

import (
	"coffee-app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user coffee.User) (int, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (phone_code, phone_hash) values ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.PhoneCode, user.Phone)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(phoneCode, phone string) (coffee.User, error) {
	var user coffee.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_code=$1 AND phone_hash=$2", userTable)
	err := r.db.Get(&user, query, phoneCode, phone)

	return user, err
}
