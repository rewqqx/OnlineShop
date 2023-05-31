package requests

import (
	"backend/src/utils/adapter"
	"backend/src/utils/database"
	"backend/src/utils/prom"
	"backend/src/validation"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type UserServer struct {
	Database *database.DBConnect
}

type UpdateUserRequest struct {
	adapter.User
	adapter.ChangePassword
}

func NewUserServer(database *database.DBConnect) *UserServer {
	return &UserServer{Database: database}
}

func (server *UserServer) GetUser(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnGetUser.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		server.GetUsers(w, r)
		//makeResponse(w, "Bad ID")
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

	w.Write([]byte(fmt.Sprintf("{\"user\" : %v}", string(json))))
}

func (server *UserServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	tokenBody := r.Header.Get("token")
	token := adapter.AuthToken{Token: tokenBody}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
	ok, err := userDatabaseAdapter.CheckTokenAndRole(token)
	if err != nil || !ok {
		makeErrorResponse(w, "bad role after auth", http.StatusBadRequest)
		return
	}

	users, err := userDatabaseAdapter.GetUsers()

	if err != nil {
		makeErrorResponse(w, "bad attempt to return users", http.StatusInternalServerError)
		return
	}

	JSON, err := json.Marshal(users)

	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"users\" : %v}", string(JSON))))
}

func (server *UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnCreateUser.Inc()
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
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

	w.Write([]byte(fmt.Sprintf("{\"token\" : %v}", string(json))))
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

	w.Write([]byte(fmt.Sprintf("{\"token\" : %v}", string(json))))
}

func (server *UserServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnUpdateUser.Inc()
	var (
		request  UpdateUserRequest
		newToken adapter.AuthToken
		JSON     []byte
	)

	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 3 {
		makeErrorResponse(w, "bad path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(dirs[2])

	tokenBody := r.Header.Get("token")
	token := adapter.AuthToken{ID: id, Token: tokenBody}

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)

	ok, err := userDatabaseAdapter.CheckToken(token)
	if err != nil || !ok {
		makeErrorResponse(w, "bad auth", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&request)
	if err != nil {
		makeErrorResponse(w, "can't parse json", http.StatusBadRequest)
		return
	}

	updateUser := request.User
	updatePassword := request.ChangePassword

	if validation.IsPasswordChangeRequest(updatePassword.Password, updatePassword.PasswordConfirmation) {
		newToken, err = userDatabaseAdapter.UpdateUserWithPassword(&updatePassword, token)
		if err != nil {
			makeErrorResponse(w, fmt.Sprintf("can not update password: %v", err), http.StatusBadRequest)
			return
		}

		JSON, err = json.Marshal(newToken)
		if err != nil {
			makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
			return
		}

	} else if validation.IsUpdateDataUserChangeRequest(updateUser.Name, updateUser.Surname, updateUser.Phone, updateUser.Mail) {
		makeErrorResponse(w, "can not set empty value to update", http.StatusBadRequest)
		return
	} else {
		newToken, err = userDatabaseAdapter.UpdateUser(&updateUser, token)
		if err != nil {
			makeErrorResponse(w, fmt.Sprintf("can not update data of user: %v", err), http.StatusBadRequest)
			return
		}

		JSON, err = json.Marshal(newToken)
		if err != nil {
			makeErrorResponse(w, "can't parse json", http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(fmt.Sprintf("{\"token\" : %v}", string(JSON))))
}
