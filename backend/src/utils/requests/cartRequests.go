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
