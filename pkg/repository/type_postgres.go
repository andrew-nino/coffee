package repository

import (
	"coffee-app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TypesPostgres struct {
	db *sqlx.DB
}

func NewTypesostgres(db *sqlx.DB) *TypesPostgres {
	return &TypesPostgres{db: db}
}

func (c *TypesPostgres) GetTypes(item string) ([]coffee.Type, error) {

	var allLists []coffee.Type

	query := fmt.Sprintf("SELECT guid, name, price FROM %s WHERE parent_guid = $1", types)
	err := c.db.Select(&allLists, query, item)

	return allLists, err
}
