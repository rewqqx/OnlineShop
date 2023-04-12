package requests

import (
	"backend/src/utils/adapter"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const USERS_COLLECTION = "users"

func GetUser(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != USERS_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		makeResponse(w, "Bad ID")
		return
	}

	tokenBody := r.Header.Get("token")
	token := adapter.AuthToken{ID: val, Token: tokenBody}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(database)
	ok, err := userDatabaseAdapter.CheckToken(token)

	if err != nil || !ok {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	user, err := userDatabaseAdapter.GetUser(token.ID)

	if err != nil {
		makeErrorResponse(w, "bad user id", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(user)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
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
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != USERS_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[1] != CREATE_ACTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	createUser := adapter.User{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&createUser)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(database)
	token, err := userDatabaseAdapter.CreateUser(&createUser)

	if err != nil {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(token)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"token\" : %v}", string(json))
	w.Write([]byte(response))
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]

	if path != "auth" {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	authData := adapter.AuthData{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&authData)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(database)
	token, err := userDatabaseAdapter.AuthUser(authData)

	if err != nil {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(token)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"status\":\"Success\", \"token\" : %v}", string(json))
	w.Write([]byte(response))
}
