package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var database *DBConnect

type StatusResponse struct {
	Status string `json:"status"`
}

func SetDatabase(connection *DBConnect) {
	database = connection
}

func setSuccessHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func makeResponse(w http.ResponseWriter, status string) error {
	response := StatusResponse{Status: status}

	jsonBody, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "{\"status\":\"Failure\"}", http.StatusBadRequest)
		return errors.New("Can't parse JSON")
	}

	w.Write(jsonBody)

	return nil
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	http.Error(w, "{\"status\":\"Success\"}", http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeResponse(w, "Bad Path")
		return
	}

	if dirs[0] != "users" {
		makeResponse(w, "Bad Path")
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		makeResponse(w, "Bad ID")
		return
	}

	var token AuthToken

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&token)

	if err != nil {
		makeResponse(w, "Bad Body")
		return
	}

	token.ID = val

	userDatabaseAdapter := UserDatabase{database: database}
	ok, err := userDatabaseAdapter.checkToken(token)

	if err != nil || !ok {
		makeResponse(w, "Bad Auth")
		return
	}

	response := "{\"status\":\"Success\"}"
	w.Write([]byte(response))
}
