package requests

import (
	"backend/src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const CREATE_ACTION = "create"

var database *utils.DBConnect

type StatusResponse struct {
	Status string `json:"status"`
}

func SetDatabase(connection *utils.DBConnect) {
	database = connection
}

func setSuccessHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func makeErrorResponse(w http.ResponseWriter, body string, status int) {
	http.Error(w, fmt.Sprintf("{\"status\":\"%v\"}", body), status)
}
func makeResponse(w http.ResponseWriter, status string) error {
	response := StatusResponse{Status: status}

	jsonBody, err := json.Marshal(response)

	if err != nil {
		makeErrorResponse(w, "Failure", http.StatusBadRequest)
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
