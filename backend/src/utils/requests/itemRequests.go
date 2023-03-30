package requests

import (
	"backend/src/utils/adapter"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const ITEM_COLLECTION = "items"

func GetItem(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeResponse(w, "Bad Path")
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeResponse(w, "Bad Path")
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		GetItems(w, r)
		return
	}

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(database)

	item, err := itemDatabaseAdapter.GetItem(val)

	if err != nil {
		makeResponse(w, "Bad Item ID")
		return
	}

	json, err := json.Marshal(item)

	if err != nil {
		makeResponse(w, "Bad JSON")
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"item\" : %v}", string(json))
	w.Write([]byte(response))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeResponse(w, "Bad Path")
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeResponse(w, "Bad Path")
		return
	}

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(database)

	pagination := adapter.Pagination{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&pagination)

	var items []*adapter.Item

	if err == nil {
		items, err = itemDatabaseAdapter.GetItemsRange(pagination)
	} else {
		items, err = itemDatabaseAdapter.GetItems()
	}

	if err != nil {
		makeResponse(w, "Bad Request")
		return
	}

	json, err := json.Marshal(items)

	if err != nil {
		makeResponse(w, "Bad JSON")
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"items\" : %v}", string(json))
	w.Write([]byte(response))
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeResponse(w, "Bad Path")
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeResponse(w, "Bad Path")
		return
	}

	if dirs[1] != CREATE_ACTION {
		makeResponse(w, "Bad Path")
		return
	}
}