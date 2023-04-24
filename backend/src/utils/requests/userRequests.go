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

type UserServer struct {
	Database *database.DBConnect
}

func NewUserServer(database *database.DBConnect) *UserServer {
	return &UserServer{Database: database}
}

const USERS_COLLECTION = "users"

func (server *UserServer) GetUser(w http.ResponseWriter, r *http.Request) {
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

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
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

func (server *UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
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

func (server *UserServer) GetToken(w http.ResponseWriter, r *http.Request) {
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

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
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

func (server *UserServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		response, token string
	)
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 3 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[0] != USERS_COLLECTION {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	if dirs[1] != "update" {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	idOfUserToUpdate := strings.Split(dirs[2], "?token")[0]
	numberIdOfUserToUpdate, err := strconv.Atoi(idOfUserToUpdate)

	updateUser := adapter.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&updateUser)
	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
	switch updateUser.Password {
	case "":
		if updateUser.Name == "" {
			makeErrorResponse(w, "can not set empty password", http.StatusBadRequest)
			return
		}

		err = userDatabaseAdapter.UpdateUser(&updateUser, numberIdOfUserToUpdate)

		if err != nil {
			makeErrorResponse(w, "can not update data of user", http.StatusBadRequest)
			return
		}

		response = fmt.Sprintf("{\"status\":\"Success\"}")
	default:
		token, err = userDatabaseAdapter.UpdateUserWithPassword(&updateUser, numberIdOfUserToUpdate)
		if err != nil {
			makeErrorResponse(w, "can not update data of user", http.StatusBadRequest)
			return
		}

		response = fmt.Sprintf("{\"status\":\"Success\", \"token\" : %v}", token)
	}

	w.Write([]byte(response))
}
