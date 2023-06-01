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

type CartServer struct {
	Database *database.Redis
}

func NewCartServer(database *database.Redis) *CartServer {
	return &CartServer{Database: database}
}

func (server *CartServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		server.SetCart(w, r)
	case http.MethodPut:
		server.AddToCart(w, r)
	case http.MethodDelete:
		server.DeleteFromCart(w, r)
	case http.MethodGet:
		server.GetCart(w, r)
	}
}

func (server *CartServer) SetCart(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	setItem := adapter.CartItem{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&setItem)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	userDatabaseAdapter := adapter.CreateCartDatabaseAdapter(server.Database)
	err = userDatabaseAdapter.SetItem(setItem)

	if err != nil {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"status\" : \"success\"}")))
}

func (server *CartServer) AddToCart(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	setItem := adapter.CartItem{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&setItem)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	userDatabaseAdapter := adapter.CreateCartDatabaseAdapter(server.Database)
	err = userDatabaseAdapter.AddItem(setItem)

	if err != nil {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"status\" : \"success\"}")))
}

func (server *CartServer) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	setItem := adapter.CartItem{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&setItem)

	userDatabaseAdapter := adapter.CreateCartDatabaseAdapter(server.Database)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if err == nil {
		err = userDatabaseAdapter.DeleteItem(setItem)
	} else {
		id, _ := strconv.Atoi(dirs[1])
		setItem.UserID = id
		err = userDatabaseAdapter.DeleteItems(setItem)
	}

	if err != nil {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"status\" : \"success\"}")))
}

func (server *CartServer) GetCart(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(dirs[1])

	if err != nil {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	itemDatabaseAdapter := adapter.CreateCartDatabaseAdapter(server.Database)

	items, err := itemDatabaseAdapter.GetCart(userID)

	if err != nil {
		makeErrorResponse(w, "bad user id", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(items)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"cart\" : %v}", string(json))))
}
