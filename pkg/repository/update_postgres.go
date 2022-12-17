package repository

import (
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

func (u *UpdatePostgres) UpdatePoints(phone string, points float32) (float32, error) {

	var outPoints float32

	createQuery := fmt.Sprintf("UPDATE %s SET value =$1 WHERE phone_hash = $2 RETURNING value", userTable)
	row := u.db.QueryRow(createQuery, points, phone)

	if err := row.Scan(&outPoints); err != nil {
		return 0, err
	}

	return outPoints, nil
}
