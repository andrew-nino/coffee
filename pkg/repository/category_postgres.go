package repository

import (
	"coffee-app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AllCategoriesPostgres struct {
	db *sqlx.DB
}

func NewAllCategoriesPostgres(db *sqlx.DB) *AllCategoriesPostgres {
	return &AllCategoriesPostgres{db: db}
}

func (c *AllCategoriesPostgres) GetCategories(category string) ([]coffee.Category, error) {

	var root = "categories"
	var sub = "sub_categories"

	var lists_root []coffee.Category
	var lists_sub []coffee.Category

	query := fmt.Sprintf("SELECT guid, name FROM %s", Categories)
	err := c.db.Select(&lists_root, query)

	query = fmt.Sprintf("SELECT guid, name FROM %s", Sub_categories)
	err = c.db.Select(&lists_sub, query)

	if root == category {

		return lists_root, err

	} else if sub == category {

		return lists_sub, err
	}

	lists_root = append(lists_root, lists_sub...)

	return lists_root, err
}
