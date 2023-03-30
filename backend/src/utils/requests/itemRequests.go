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
		makeResponse(w, "Bad ID")
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
