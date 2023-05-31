package requests

import (
	"backend/src/utils/prom"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func setSuccessHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("Content-Type", "application/json")
}

func makeErrorResponse(w http.ResponseWriter, body string, status int) {
	w.WriteHeader(status)
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
	w.WriteHeader(http.StatusOK)

	return nil
}

func Ping(w http.ResponseWriter, r *http.Request) {
	prom.MetricOnPing.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	http.Error(w, "{\"status\":\"Success\"}", http.StatusOK)
}
