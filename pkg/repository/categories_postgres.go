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

func (c *AllCategoriesPostgres) GetALLCategories() ([]coffee.Categories, error) {
	var allLists []coffee.Categories
	var allLists_2 []coffee.Categories

	query := fmt.Sprintf("SELECT guid, name FROM %s", Categories)
	err := c.db.Select(&allLists, query)

	query  = fmt.Sprintf("SELECT guid, name FROM %s", Sub_categories)
	err = c.db.Select(&allLists_2, query)

	allLists = append(allLists, allLists_2...)

	return allLists, err
}
