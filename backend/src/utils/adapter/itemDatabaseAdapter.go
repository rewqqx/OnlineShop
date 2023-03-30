package adapter

import (
	"backend/src/utils"
	"fmt"
)

type ItemDatabase struct {
	database *utils.DBConnect
}

const ITEMS_TABLE_NAME = "items"

type Item struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"item_name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
	ImageIDS    []int  `json:"image_ids" db:"image_ids"`
}

func CreateItemDatabaseAdapter(database *utils.DBConnect) *ItemDatabase {
	adapter := &ItemDatabase{database: database}
	return adapter
}

func (adapter *ItemDatabase) GetItems(count int, page int) (items []*Item, err error) {
	rows, err := adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v", ITEMS_TABLE_NAME))

	for rows.Next() {
		item := &Item{}
		err = rows.Scan(item)
		items = append(items, item)
	}

	return
}

func (adapter *ItemDatabase) GetItem(id int) (item *Item, err error) {
	item = &Item{}
	err = adapter.database.Connection.Get(item, fmt.Sprintf("SELECT * FROM online_shop.%v WHERE id=$1", ITEMS_TABLE_NAME), id)

	return
}

func (adapter *ItemDatabase) CreateItem(item *Item) (id int64, err error) {
	res, err := adapter.database.Connection.Exec(fmt.Sprintf("INSERT INTO online_shop.%v (iterm_name,price,description,image_ids) VALUES ($1, $2, $3, $4)", ITEMS_TABLE_NAME), item.Name, item.Price, item.Description, item.ImageIDS)
	return res.LastInsertId()
}

func (adapter *ItemDatabase) DeleteItem(id int) (err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("DELETE FROM online_shop.%v WHERE id=$1", ITEMS_TABLE_NAME), id)
	return
}
