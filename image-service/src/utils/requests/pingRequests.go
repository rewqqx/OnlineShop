package requests

import (
	"fmt"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func setSuccessHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func makeErrorResponse(w http.ResponseWriter, body string, status int) {
	http.Error(w, fmt.Sprintf("{\"status\":\"%v\"}", body), status)
	w.WriteHeader(status)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	http.Error(w, "{\"status\":\"Success\"}", http.StatusOK)
}
