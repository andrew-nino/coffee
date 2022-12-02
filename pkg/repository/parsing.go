package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
)

type YTime struct {
	Success bool   `json:"success"`
	Count   int    `json:"-"`
	Rows    []List `json:"rows"`
	Error   string `json:"-"`
}

type List struct {
	Guid         string      `json:"guid"`
	Name         string      `json:"name"`
	Priority     int         `json:"priority"`
	ImageLink    string      `json:"imageLink"`
	CategoryList []List      `json:"categoryList"`
	ItemLists    []ItemsList `json:"itemList"`
	GoodsLists   []GoodList  `json:"goodsList"`
	ComboList    string      `json:"combolist,omitempty"`
}

type ItemsList struct {
	Guid         string     `json:"guid"`
	Name         string     `json:"name"`
	Priority     int        `json:"priority,omitempty"`
	Price        float32    `json:"price,omitempty"`
	ImageLink    string     `json:"imageLink,omitempty"`
	Desckription string     `json:"description,omitempty"`
	Recipe       string     `json:"recipe,omitempty"`
	TypeLists    []TypeList `json:"typeList"`
	// SupplementCategoryToFreeCount map[string]int
	// DefaultSupplements            []string
}

type TypeList struct {
	Guid   string  `json:"guid,omitempty"`
	Name   string  `json:"name,omitempty"`
	Price  float32 `json:"price,omitempty"`
	IsTogo bool    `json:"isTogo,omitempty"`
}

type GoodList struct {
	Guid         string  `json:"guid,omitempty"`
	Name         string  `json:"name,omitempty"`
	Priority     int     `json:"priority,omitempty"`
	Price        float32 `json:"price,omitempty"`
	ImageLink    string  `json:"imageLink,omitempty"`
	Desckription string  `json:"description,omitempty"`
	Recipe       string  `json:"recipe,omitempty"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parsingDB(tx *sqlx.DB) {

	file, err := ioutil.ReadFile("/home/andrew/Work/Go/src/coffee/response.json")
	check(err)

	var newStruct YTime

	err = json.Unmarshal(file, &newStruct)
	check(err)

	for _, rows := range newStruct.Rows {

		createQuery := fmt.Sprintf("INSERT INTO %s (guid, name) VALUES ($1, $2)", categories)
		_, err := tx.Exec(createQuery, rows.Guid, rows.Name)
		check(err)

		for _, categoryList := range rows.CategoryList {

			createQuery := fmt.Sprintf("INSERT INTO %s (parent_guid, guid, name) VALUES ($1, $2, $3)", sub_categories)
			_, err := tx.Exec(createQuery, rows.Guid, categoryList.Guid, categoryList.Name)
			check(err)

			for _, item := range categoryList.ItemLists {

				createQuery := fmt.Sprintf("INSERT INTO %s (cat_guid, sub_cat_guid, guid, name, description) VALUES ($1, $2, $3, $4,$5)", items)
				_, err := tx.Exec(createQuery, rows.Guid, categoryList.Guid, item.Guid, item.Name, item.Desckription)
				check(err)

				for _, typeList := range item.TypeLists {

					createQuery := fmt.Sprintf("INSERT INTO %s (parent_guid, guid, name, price) VALUES ($1, $2, $3, $4)", types)
					_, err := tx.Exec(createQuery, item.Guid, typeList.Guid, typeList.Name, typeList.Price)
					check(err)
				}
				// for _, supplement := range items.SupplementCategoryToFreeCount {
				// 	for k, v := range supplement {
				// 		fmt.Printf("%s:%d", k, v)
				// 	}
				// }
				// for _, defaultSupplement := range items.DefaultSupplements {
				// 	fmt.Println(defaultSupplement)
				// }
			}
			// for _, goodList := range categoryList.GoodsLists {
			// }

			// fmt.Println(categoryList.ComboList)
		}
		for _, itemsList := range rows.ItemLists {

			createQuery := fmt.Sprintf("INSERT INTO %s (cat_guid, guid, name, description) VALUES ($1, $2, $3, $4)", items)
			_, err := tx.Exec(createQuery, rows.Guid, itemsList.Guid, itemsList.Name, itemsList.Desckription)
			check(err)

			for _, typeList := range itemsList.TypeLists {

				createQuery := fmt.Sprintf("INSERT INTO %s (parent_guid, guid, name, price) VALUES ($1, $2, $3, $4)", types)
				_, err := tx.Exec(createQuery, itemsList.Guid, typeList.Guid, typeList.Name, typeList.Price)
				check(err)
			}
			// for _, supplement := range itemsList.SupplementCategoryToFreeCount {
			// 	for k, v := range supplement {
			// 		fmt.Printf("%s:%d", k, v)
			// 	}
			// }
			// for _, defaultSupplement := range itemsList.DefaultSupplements {
			// 	fmt.Println(defaultSupplement)
			// }
		}
		// for _, goodList := range rows.GoodsLists {
		// 	fmt.Println(goodList.Guid)
		// 	fmt.Println(goodList.Name)
		// 	fmt.Println(goodList.Price)
		// 	fmt.Println(goodList.Desckription)
		// }

		// fmt.Println(rows.ComboList)
	}
}