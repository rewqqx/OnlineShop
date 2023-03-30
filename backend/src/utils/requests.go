package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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

	tokenBody := r.Header.Get("token")
	token := AuthToken{ID: val, Token: tokenBody}

	userDatabaseAdapter := UserDatabase{database: database}
	ok, err := userDatabaseAdapter.checkToken(token)

	if err != nil || !ok {
		makeResponse(w, "Bad Auth")
		return
	}

	user, err := userDatabaseAdapter.getUser(token.ID)

	if err != nil {
		makeResponse(w, "Bad User ID")
		return
	}

	json, err := json.Marshal(user)

	if err != nil {
		makeResponse(w, "Bad JSON")
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"user\" : %v}", string(json))
	w.Write([]byte(response))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if dirs[1] != "create" {
		makeResponse(w, "Bad Path")
		return
	}

	createUser := User{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&createUser)

	if err != nil {
		makeResponse(w, "Bad Body")
		return
	}

	userDatabaseAdapter := UserDatabase{database: database}
	token, err := userDatabaseAdapter.createUser(&createUser)

	if err != nil {
		makeResponse(w, "Bad Auth")
		return
	}

	json, err := json.Marshal(token)

	if err != nil {
		makeResponse(w, "Bad JSON")
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"token\" : %v}", string(json))
	w.Write([]byte(response))
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]

	if path != "auth" {
		makeResponse(w, "Bad Path")
		return
	}

	authData := AuthData{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&authData)

	if err != nil {
		makeResponse(w, "Bad Body")
		return
	}

	userDatabaseAdapter := UserDatabase{database: database}
	token, err := userDatabaseAdapter.authUser(authData)

	if err != nil {
		makeResponse(w, "Bad Auth")
		return
	}

	json, err := json.Marshal(token)

	if err != nil {
		makeResponse(w, "Bad JSON")
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"token\" : %v}", string(json))
	w.Write([]byte(response))
}
