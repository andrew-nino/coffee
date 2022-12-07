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
	var allItems []coffee.Item

	query := fmt.Sprintf("SELECT guid, name,description FROM %s WHERE (cat_guid = $1 AND sub_cat_guid = '') OR (sub_cat_guid = $1 AND cat_guid != $1)", items)
	err := c.db.Select(&allItems, query, categoty)

	for i := 0; i < len(allItems); i++ {

		var allTypes []coffee.Type

		query = fmt.Sprintf("SELECT name, price FROM %s WHERE parent_guid = $1", types)
		_ = c.db.Select(&allTypes, query, allItems[i].Guid)

		allItems[i].Types = append(allItems[i].Types, allTypes...)
	}

	return allItems, err
}
