package repository

import (
	"coffee-app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ActionPostgres struct {
	db *sqlx.DB
}

func NewActionPostgres(db *sqlx.DB) *ActionPostgres {
	return &ActionPostgres{db: db}
}

func (c *ActionPostgres) GetActions() ([]coffee.Action, error) {
	var allActions []coffee.Action

	query := fmt.Sprintf("SELECT guid, name, start_date, expiry_date, picture FROM %s ", actions)
	err := c.db.Select(&allActions, query)

	return allActions, err
}

func (c *ActionPostgres) GetActionById(guid string) (coffee.Action, error) {
	var action coffee.Action

	query := fmt.Sprintf("SELECT guid, name, start_date, expiry_date, picture, description FROM %s WHERE guid = $1", actions)
	err := c.db.Get(&action, query, guid)

	return action, err
}
