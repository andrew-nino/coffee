package repository

import (
	"coffee-app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ItemsPostgres struct {
	db *sqlx.DB
}

func NewItemsostgres(db *sqlx.DB) *ItemsPostgres {
	return &ItemsPostgres{db: db}
}

func (c *ItemsPostgres) GetItemsById(categoty string) ([]coffee.Items, error) {
	var allLists []coffee.Items

	query := fmt.Sprintf("SELECT guid, name FROM %s WHERE cat_guid = $1 or sub_cat_guid = $1", Items)
	err := c.db.Select(&allLists, query, categoty)

	return allLists, err
}
