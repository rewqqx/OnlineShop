package utils

import (
	"backend/src/utils/adapter"
	"backend/src/utils/database"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

const HOST = "127.0.0.1"
const contentType = "application/json"
const successStatus = "{\"status\":\"Success\"}\n"

type Response struct {
	Status string         `json:"status"`
	Items  []adapter.Item `json:"items"`
}

type Create struct {
	Status string `json:"status"`
	Token  struct {
		ID    int    `json:"ID"`
		Token string `json:"token"`
	} `json:"token"`
}

func startingServer() (err error, databaseConnection database.DBConnect) {
	databaseConnection = database.DBConnect{Ip: HOST, Port: "5432", Password: "pgpass", User: "postgres", Database: "postgres"}
	err = databaseConnection.Open()
	server := New(&databaseConnection)
	go server.Start(8080)
	return
}

func ReadBody(respCreateUser *http.Response) (bodyString string, err error) {
	bodyBytes, err := io.ReadAll(respCreateUser.Body)
	bodyString = string(bodyBytes)
	return
}

func TestPing(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	resp, err := http.Get("http://" + HOST + ":8080/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)
	require.Equal(t, successStatus, bodyString, "Error in body of response. Expected %s actually %s",
		successStatus, bodyString)
}

func TestGetItem(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	resp, err := http.Get("http://" + HOST + ":8080/items/1/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)

	item := &adapter.Item{}
	err = db.Connection.Get(item, fmt.Sprintf("SELECT * FROM online_shop.items WHERE id=1"))
	JSON, err := json.Marshal(item)
	expected := fmt.Sprintf("{\"status\":\"Success\", \"item\" : %v}", string(JSON))
	require.Equal(t, expected, bodyString)
}

func TestGetItemsSuccess(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(adapter.Pagination{Offset: 0, Limit: 5})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+HOST+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)

	var response Response
	err = json.Unmarshal([]byte(bodyString), &response)

	require.Equal(t, nil, err, "Error unmarshalling JSON:", err)
	require.Equal(t, 5, len(response.Items), "Error in pagination. Expected 5, actually %d",
		len(response.Items))
}

func TestGetItemsSuccessOneItem(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(adapter.Pagination{Offset: 1, Limit: 1})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+HOST+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)

	var response Response
	err = json.Unmarshal([]byte(bodyString), &response)

	require.Equal(t, nil, err, "Error unmarshalling JSON:", err)
	require.Equal(t, 1, len(response.Items), "Error in pagination. Expected 5, actually %d",
		len(response.Items))
}

func TestGetItemsUnsuccessful(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(adapter.Pagination{Offset: 0, Limit: -1})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+HOST+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Error in request: %v", http.StatusOK)
}

func TestCreateUserSuccess(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{ID: -1, Name: "Bogdan", Surname: "Madzhuga", Patronymic: "Andreevich", Phone: "", Birthdate: nil, Mail: "madzhuga@mail.ru", Password: "bogdan0308", RoleId: 2, Token: "", Sex: 1})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+HOST+":8080/users/create", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)
}

func TestCreateUserSuccessAndAuthSuccess(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{ID: -1, Name: "Bogdan", Surname: "Madzhuga", Patronymic: "Andreevich", Phone: "", Birthdate: nil, Mail: "madzhuga@mail.ru", Password: "bogdan0308", RoleId: 2, Token: "", Sex: 1})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	respCreateUser, err := http.Post("http://"+HOST+":8080/users/create", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respCreateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, respCreateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respCreateUser)

	var create Create
	err = json.Unmarshal([]byte(bodyString), &create)

	respAuth, err := http.Get("http://" + HOST + ":8080/?token=" + create.Token.Token)
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respAuth.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, respAuth.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyStringOfAuth, err := ReadBody(respAuth)
	require.Equal(t, successStatus, bodyStringOfAuth, "Error in body of response. Expected %s actually %s",
		successStatus, bodyStringOfAuth)
}
