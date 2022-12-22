package repository

import (
	"coffee-app"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type UpdatePostgres struct {
	db *sqlx.DB
}

func NewUpdatePostgres(db *sqlx.DB) *UpdatePostgres {
	return &UpdatePostgres{db: db}
}

func (u *UpdatePostgres) UpdateDB() (string, error) {

	parsingDB(u.db)

	return time.Now().GoString(), nil
}

func (u *UpdatePostgres) UpdatePoints(phone string, points float32) (coffee.User, error) {

	var user coffee.User

	createQuery := fmt.Sprintf("UPDATE %s SET value =$1 WHERE phone_hash = $2 RETURNING value, message_key", userTable)
	err := u.db.Get(&user, createQuery, points, phone)

	return user, err
}

func (u *UpdatePostgres) UpdateUser(user coffee.User) error {

	// if user.Name != ""
	// if user.Surname != ""
	// if user.Birthday != ""

	var key string

	createQuery := fmt.Sprintf("SELECT message_key FROM %s WHERE phone_hash = $1", userTable)
	err := u.db.Get(&key, createQuery, user.Phone)
	if err != nil {
		return err
	}

	if user.MessageKey != key {
		createQuery := fmt.Sprintf("UPDATE %s SET message_key=$1 WHERE phone_hash = $2 ", userTable)
		_, err := u.db.Exec(createQuery, user.MessageKey, user.Phone)
		if err != nil {
			return err
		}
	}

	return nil
}
