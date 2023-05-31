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

type TagServer struct {
	Database *database.DBConnect
}

func NewTagServer(database *database.DBConnect) *TagServer {
	return &TagServer{Database: database}
}

func (server *TagServer) GetTag(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnGETTegs.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		server.GetTags(w, r)
		return
	}

	tagDatabaseAdapter := adapter.CreateTagDatabaseAdapter(server.Database)

	tag, err := tagDatabaseAdapter.GetTag(val)

	if err != nil {
		makeErrorResponse(w, "bad tag id", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(tag)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"tag\" : %v}", string(json))))
}

func (server *TagServer) GetTags(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnGETTegs.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	tagDatabaseAdapter := adapter.CreateTagDatabaseAdapter(server.Database)

	tags, err := tagDatabaseAdapter.GetTags()

	if err != nil {
		makeErrorResponse(w, "can't find tags", http.StatusInternalServerError)
		return
	}

	JSON, err := json.Marshal(tags)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"tags\" : %v}", string(JSON))))
}
