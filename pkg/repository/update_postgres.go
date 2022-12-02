package repository

import (
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
