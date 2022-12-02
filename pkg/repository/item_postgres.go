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

func (c *ItemsPostgres) GetItems() ([]coffee.Item, error) {
	var allLists []coffee.Item

	query := fmt.Sprintf("SELECT guid, name, description FROM %s ", items)
	err := c.db.Select(&allLists, query)

	return allLists, err
}

func (c *ItemsPostgres) GetItemsById(categoty string) ([]coffee.Item, error) {
	var allLists []coffee.Item

	query := fmt.Sprintf("SELECT guid, name,description FROM %s WHERE (cat_guid = $1 AND sub_cat_guid = '') OR (sub_cat_guid = $1 AND cat_guid != $1)", items)
	err := c.db.Select(&allLists, query, categoty)

	return allLists, err
}