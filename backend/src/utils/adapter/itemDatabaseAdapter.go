package adapter

import (
	"backend/src/utils/database"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type ItemDatabase struct {
	database *database.DBConnect
}

const ITEMS_TABLE_NAME = "items"

type Item struct {
	ID          int           `json:"id" db:"id"`
	Name        string        `json:"name" db:"item_name"`
	Price       int           `json:"price" db:"price"`
	Description string        `json:"description" db:"description"`
	ImageIDS    pq.Int64Array `json:"image_ids" db:"image_ids"`
	TagIDS      pq.Int64Array `json:"tag_ids" db:"tag_ids"`
}

type ItemFilters struct {
	Offset *int `json:"offset" db:"offset"`
	Limit  *int `json:"limit" db:"limit"`
	Tag    *int `json:"tag" db:"tag"`
}

func CreateItemDatabaseAdapter(database *database.DBConnect) *ItemDatabase {
	adapter := &ItemDatabase{database: database}
	return adapter
}

func (adapter *ItemDatabase) GetItems() (items []*Item, err error) {
	rows, err := adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v", ITEMS_TABLE_NAME))

	if err != nil {
		return
	}

	return parseItemsFromRows(rows)
}

func (adapter *ItemDatabase) GetItemsRange(filters ItemFilters) (items []*Item, err error) {

	var rows *sql.Rows

	offset := 0

	if filters.Offset != nil {
		offset = *filters.Offset
	}

	if filters.Limit != nil && filters.Tag == nil {
		rows, err = adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY;", ITEMS_TABLE_NAME), offset, *filters.Limit)
	} else if filters.Limit == nil && filters.Tag != nil {
		rows, err = adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v WHERE $1 = ANY(tag_ids);", ITEMS_TABLE_NAME), *filters.Tag)
	} else if filters.Limit != nil && filters.Tag != nil {
		rows, err = adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v WHERE $1 = ANY(tag_ids) OFFSET $2 ROWS FETCH NEXT $3 ROWS ONLY;", ITEMS_TABLE_NAME), *filters.Tag, offset, *filters.Limit)
	}

	if err != nil {
		return
	}

	return parseItemsFromRows(rows)
}

func parseItemsFromRows(rows *sql.Rows) (items []*Item, err error) {
	for rows.Next() {
		var id int
		var itemName string
		var price int
		var desc string
		var imageIds pq.Int64Array
		var tagIds pq.Int64Array

		err = rows.Scan(&id, &itemName, &price, &desc, &imageIds, &tagIds)

		if err != nil {
			return
		}

		item := &Item{id, itemName, price, desc, imageIds, tagIds}
		items = append(items, item)
	}

	return
}

func (adapter *ItemDatabase) GetItem(id int) (item *Item, err error) {
	item = &Item{}
	err = adapter.database.Connection.Get(item, fmt.Sprintf("SELECT * FROM online_shop.%v WHERE id=$1", ITEMS_TABLE_NAME), id)

	return
}

func (adapter *ItemDatabase) CreateItem(item *Item) (int64, error) {
	res, _ := adapter.database.Connection.Exec(fmt.Sprintf("INSERT INTO online_shop.%v (item_name,price,description,image_ids,tag_ids) VALUES ($1, $2, $3, $4, $5)", ITEMS_TABLE_NAME), item.Name, item.Price, item.Description, item.ImageIDS, item.TagIDS)
	//return
	return res.LastInsertId()
}

func (adapter *ItemDatabase) DeleteItem(id int) (err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("DELETE FROM online_shop.%v WHERE id=$1", ITEMS_TABLE_NAME), id)
	return
}

func (adapter *ItemDatabase) UpdateItem(item *Item, id int) (err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET item_name = $1, price = $2, description = $3, image_ids = $4, tag_ids = $5 WHERE id = $6", ITEMS_TABLE_NAME), item.Name, item.Price, item.Description, item.ImageIDS, item.TagIDS, id)
	return
}
