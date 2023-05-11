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
	"strconv"
	"testing"
)

const HOST = "127.0.0.1"
const contentTypeJSON = "application/json"
const contentTypeText = "text/plain; charset=utf-8"
const IDUserForUpdate = 1

type Response struct {
	Status string         `json:"status"`
	Items  []adapter.Item `json:"items"`
}

//host := "127.0.0.1"
//database := database.DBConnect{Ip: host, Port: "6000", Password: "pgpass", User: "postgres", Database: "postgres"}
//err := database.Open()

type Create struct {
	Token struct {
		ID    int    `json:"ID" db:"id"`
		Token string `json:"token" db:"token"`
	} `json:"token"`
}

type StatusOfResponse struct {
	Status string `json:"status"`
}

func startingServer() (err error, databaseConnection database.DBConnect) {
	databaseConnection = database.DBConnect{Ip: HOST, Port: "6000", Password: "pgpass", User: "postgres", Database: "postgres"}
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

func getTokenWIthExpected(db database.DBConnect, id int) (token *adapter.AuthToken, expected string) {
	token = &adapter.AuthToken{}
	db.Connection.Get(token, fmt.Sprintf("SELECT id, token FROM online_shop.users WHERE id = $1"), id)
	JSON, _ := json.Marshal(token)
	expected = fmt.Sprintf("{\"token\" : %v}", string(JSON))
	return
}

func makeStructOutputForTest(bodyString, mistake string) (actualOutput, expected StatusOfResponse) {
	var actual StatusOfResponse
	_ = json.Unmarshal([]byte(bodyString), &actual)

	expected = StatusOfResponse{}
	expected.Status = mistake
	return actual, expected
}

func TestPing(t *testing.T) {
	err, _ := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	resp, err := http.Get("http://" + HOST + ":8080/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeJSON, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)
	actual, expected := makeStructOutputForTest(bodyString, "Success")
	require.Equal(t, expected, actual, "Error in body of response. Expected %s actually %s",
		expected, actual)
}

func TestGetItem(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	resp, err := http.Get("http://" + HOST + ":8080/items/1/")

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeJSON, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(resp)

	item := &adapter.Item{}
	err = db.Connection.Get(item, fmt.Sprintf("SELECT * FROM online_shop.items WHERE id=1"))
	JSON, err := json.Marshal(item)
	expected := fmt.Sprintf("{\"item\" : %v}", string(JSON))
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
	require.Equal(t, contentTypeJSON, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
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
	require.Equal(t, contentTypeJSON, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
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
	require.Equal(t, contentTypeJSON, resp.Header.Get("Content-Type"), "Error in header Content-Type: %v",
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
	require.Equal(t, contentTypeJSON, respCreateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respCreateUser)

	var create Create
	err = json.Unmarshal([]byte(bodyString), &create)

	respAuth, err := http.Get("http://" + HOST + ":8080/?token=" + create.Token.Token)
	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respAuth.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeJSON, respAuth.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyStringOfAuth, err := ReadBody(respAuth)

	var actual StatusOfResponse
	err = json.Unmarshal([]byte(bodyStringOfAuth), &actual)

	expected := StatusOfResponse{}
	expected.Status = "Success"

	require.Equal(t, expected, actual, "Error in body of response. Expected %s actually %s",
		expected, actual)
}

func TestUpdateUserDataSuccess(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{Name: "huhu4r", Surname: "v4TMа54вавваваadfdzhuga", Phone: "+79843435", Mail: "v1tdfdfrdzhfrd3a@mail.ru"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, expected := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeJSON, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, expected, bodyString, "Error in response body. Expected: %v, actually %v", expected, bodyString)
}

func TestUpdateUserDataInvalidName(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{Name: "hu", Surname: "v4TMа54вавваваadfdzhuga", Phone: "+79843435", Mail: "v1tdfdfrdzhfrd3a@mail.ru"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update data of user: name must contains more than 3 symbols")
	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserDataInvalidSurname(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{Name: "hutttff", Surname: "v", Phone: "+79843435", Mail: "v1tdfdfrdzhfrd3a@mail.ru"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update data of user: surname must contains more than 3 symbols")
	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserDataInvalidPhone(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{Name: "hutttff", Surname: "vfdfdfdf", Phone: "79843435", Mail: "v1tdfdfrdzhfrd3a@mail.ru"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update data of user: phone must contains +")
	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserDataInvalidMail(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.User{Name: "hutttff", Surname: "vfdfdfdf", Phone: "+5579843435", Mail: "v1tdfdfrdzhfrd3amail.ru"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update data of user: mail must contains @")
	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserPasswordInvalidMatch(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.ChangePassword{Password: "qwerty1234", PasswordConfirmation: "qwerty123"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update password: passwords must match")
	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserPasswordInvalidPassword(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.ChangePassword{Password: "q", PasswordConfirmation: "q"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, _ := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusBadRequest, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeText, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.Equal(t, nil, err, "Error in reading body: %v", err)

	actual, expected := makeStructOutputForTest(bodyString, "can not update password: password must contains more than 7 symbols")

	require.Equal(t, expected, actual, "Error in request: %v", err)
}

func TestUpdateUserPasswordSuccess(t *testing.T) {
	err, db := startingServer()
	require.Equal(t, nil, err, "Error in DB connection: %v", err)

	jsonPayload, err := json.Marshal(&adapter.ChangePassword{Password: "qwerty12", PasswordConfirmation: "qwerty12"})
	require.Equal(t, nil, err, "Error in Marshal: %v", err)

	token, expected := getTokenWIthExpected(db, IDUserForUpdate)

	req, err := http.NewRequest("POST", "http://"+HOST+":8080/users/update/"+strconv.Itoa(IDUserForUpdate), bytes.NewBuffer(jsonPayload))
	require.Equal(t, nil, err, "Error in making request: %v", err)
	req.Header.Set("token", token.Token)

	respUpdateUser, err := http.DefaultClient.Do(req)

	require.Equal(t, nil, err, "Error in request: %v", err)
	require.Equal(t, http.StatusOK, respUpdateUser.StatusCode, "Error in request: %v", http.StatusOK)
	require.Equal(t, contentTypeJSON, respUpdateUser.Header.Get("Content-Type"), "Error in header Content-Type: %v",
		err)

	bodyString, err := ReadBody(respUpdateUser)
	require.NotEqual(t, expected, bodyString)
}
