package utils

import (
	"backend/src/utils/database"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

const contentType = "application/json"
const successStatus = "{\"status\":\"Success\"}\n"
const JSONExample = "{\"status\":\"Success\", \"item\" : {\"id\":1,\"name\":\"Apple\",\"price\":1,\"description\":\"" +
	"Sweee Apple\",\"image_ids\":[1,2]}}"

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	ImageIDs    []int  `json:"image_ids"`
}

type Response struct {
	Status string `json:"status"`
	Items  []Item `json:"items"`
}
type Create struct {
	Status string `json:"status"`
	Token  struct {
		ID    int    `json:"ID"`
		Token string `json:"token"`
	} `json:"token"`
}

func TestPing(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	resp, err := http.Get("http://" + host + ":8080/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	require.Equal(t, successStatus, bodyString, "Error in body of response. Expected %s actually %s",
		successStatus, bodyString)
}

func TestGetItem(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	resp, err := http.Get("http://" + host + ":8080/items/1/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	require.Equal(t, JSONExample, bodyString)
}

func TestGetItemsSuccess(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	payload := map[string]int{
		"offset": 0,
		"limit":  5,
	}

	jsonPayload, err := json.Marshal(payload)
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+host+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	var response Response
	err = json.Unmarshal([]byte(bodyString), &response)

	require.Equal(t, nil, err, "Error unmarshalling JSON:", err)
	require.Equal(t, 5, len(response.Items), "Error in pagination. Expected 5, actually %d",
		len(response.Items))
}

func TestGetItemsSuccessOneItem(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	payload := map[string]int{
		"offset": 1,
		"limit":  1,
	}

	jsonPayload, err := json.Marshal(payload)
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+host+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	var response Response
	err = json.Unmarshal([]byte(bodyString), &response)

	require.Equal(t, nil, err, "Error unmarshalling JSON:", err)
	require.Equal(t, 1, len(response.Items), "Error in pagination. Expected 5, actually %d",
		len(response.Items))
}

func TestGetItemsUnsuccessful(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	payload := map[string]int{
		"offset": 0,
		"limit":  -1,
	}

	jsonPayload, err := json.Marshal(payload)
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+host+":8080/items/", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Error in request: %v", http.StatusOK)
}

func TestCreateUserSuccess(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	payload := map[string]any{
		"id":              -1,
		"user_name":       "Bogdan",
		"user_surname":    "Madzhuga",
		"user_patronymic": "Andreevich",
		"phone":           "",
		"birthdate":       nil,
		"mail":            "madzhuga@mail.ru",
		"password_hash":   "bogdan0308",
		"role_id":         2,
		"token":           "",
	}

	jsonPayload, err := json.Marshal(payload)
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	resp, err := http.Post("http://"+host+":8080/users/create", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)
}

func TestCreateUserSuccessAndAuthSuccess(t *testing.T) {

	host := "127.0.0.1"
	databaseConnection := database.DBConnect{Ip: host, Port: "5433", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := databaseConnection.Open()

	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	server := New(&databaseConnection)
	go server.Start(8080)

	payload := map[string]any{
		"id":              -1,
		"user_name":       "Bogdan",
		"user_surname":    "Madzhuga",
		"user_patronymic": "Andreevich",
		"phone":           "",
		"birthdate":       nil,
		"mail":            "madzhuga@mail.ru",
		"password_hash":   "bogdan0308",
		"role_id":         2,
		"token":           "",
	}

	jsonPayload, err := json.Marshal(payload)
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	respCreateUser, err := http.Post("http://"+host+":8080/users/create", "application/json", bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respCreateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, respCreateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytes, err := io.ReadAll(respCreateUser.Body)
	bodyString := string(bodyBytes)

	var create Create
	err = json.Unmarshal([]byte(bodyString), &create)

	respAuth, err := http.Get("http://" + host + ":8080/?token=" + create.Token.Token)
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respAuth.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentType, respAuth.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyBytesOfAuth, err := io.ReadAll(respAuth.Body)
	bodyStringOfAuth := string(bodyBytesOfAuth)
	require.Equal(t, successStatus, bodyStringOfAuth, "Error in body of response. Expected %s actually %s",
		successStatus, bodyStringOfAuth)
}
