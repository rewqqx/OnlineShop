package requests

import (
	"backend/src/utils/adapter"
	"backend/src/utils/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const ITEM_COLLECTION = "items"

type ItemServer struct {
	Database *database.DBConnect
}

func NewItemServer(database *database.DBConnect) *ItemServer {
	return &ItemServer{Database: database}
}

func (server *ItemServer) GetItem(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		server.GetItems(w, r)
		return
	}

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(server.Database)

	item, err := itemDatabaseAdapter.GetItem(val)

	if err != nil {
		makeErrorResponse(w, "bad item id", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(item)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"item\" : %v}", string(json))
	w.Write([]byte(response))
}

func (server *ItemServer) GetItems(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(server.Database)

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
		makeErrorResponse(w, "can't find items", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(items)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"items\" : %v}", string(json))
	w.Write([]byte(response))
}

func (server *ItemServer) CreateItem(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != ITEM_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[1] != CREATE_ACTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}
}
