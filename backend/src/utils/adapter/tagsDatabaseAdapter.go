package adapter

import (
	"backend/src/utils/database"
	"database/sql"
	"fmt"
)

const TAGS_TABLE_NAME = "tags"

type TagDatabase struct {
	database *database.DBConnect
}

func CreateTagDatabaseAdapter(database *database.DBConnect) *TagDatabase {
	adapter := &TagDatabase{database: database}
	return adapter
}

type Tag struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"tag_name" db:"tag_name"`
	ParentID int    `json:"parent_id" db:"parent_id"`
}

func (adapter *TagDatabase) GetTag(id int) (tag *Tag, err error) {
	tag = &Tag{}
	err = adapter.database.Connection.Get(tag, fmt.Sprintf("SELECT * FROM online_shop.%v WHERE id=$1", TAGS_TABLE_NAME), id)

	return
}

func (adapter *TagDatabase) GetTags() (tags []*Tag, err error) {
	rows, err := adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v", TAGS_TABLE_NAME))

	if err != nil {
		return
	}

	return parseTagsFromRows(rows)
}

func parseTagsFromRows(rows *sql.Rows) (tags []*Tag, err error) {
	for rows.Next() {
		var id int
		var tagName string
		var parentID int

		err = rows.Scan(&id, &tagName, &parentID)

		if err != nil {
			return
		}

		tag := &Tag{id, tagName, parentID}
		tags = append(tags, tag)
	}

	return
}
