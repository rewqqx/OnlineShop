package adapter

import "backend/src/utils"

type ItemDatabase struct {
	database *utils.DBConnect
}

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
