package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) GetBalance(id int) (float32, error) {

	var value float32

	query := fmt.Sprintf("SELECT value FROM %s WHERE id = $1", userTable)
	err := c.db.Get(&value, query, id)

	return value, err
}
