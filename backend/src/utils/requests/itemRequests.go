package requests

import (
	"backend/src/utils/adapter"
	"backend/src/utils/database"
	"backend/src/utils/prom"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ItemServer struct {
	Database *database.DBConnect
}

func NewItemServer(database *database.DBConnect) *ItemServer {
	return &ItemServer{Database: database}
}

func (server *ItemServer) GetItem(w http.ResponseWriter, r *http.Request) {
	//ready
	prom.MetricOnGETItems.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
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

	w.Write([]byte(fmt.Sprintf("{\"item\" : %v}", string(json))))
}

func (server *ItemServer) GetItems(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnGETItems.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(server.Database)

	filters := adapter.ItemFilters{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&filters)

	var items []*adapter.Item

	if err == nil {
		items, err = itemDatabaseAdapter.GetItemsRange(filters)
	} else {
		items, err = itemDatabaseAdapter.GetItems()
	}

	if err != nil {
		makeErrorResponse(w, "can't find items", http.StatusInternalServerError)
		return
	}

	JSON, err := json.Marshal(items)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"items\" : %v}", string(JSON))))
}

func (server *ItemServer) CreateItem(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnCreateItems.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}
}

func (server *ItemServer) DeleteItem(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(dirs[2])

	itemDatabaseAdapter := adapter.CreateItemDatabaseAdapter(server.Database)

	err := itemDatabaseAdapter.DeleteItem(id)
	if err != nil {
		makeErrorResponse(w, "can't find item", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"status\" : \"success\"}")))
}
