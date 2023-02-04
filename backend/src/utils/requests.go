package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type StatusResponse struct {
	Status string `json:"status"`
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
		http.Error(w, "{'status':'failure'}", http.StatusBadRequest)
		return errors.New("Can't parse JSON")
	}

	w.Write(jsonBody)

	return nil
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	http.Error(w, "{'status':'success'}", http.StatusOK)
}

func TablesList(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")

	if len(dirs) < 2 {
		makeResponse(w, "failure")
		return
	}

	if dirs[0] != "tables" {
		makeResponse(w, "wrong path")
		return
	}

	parentSchema := dirs[1]

	var connect DBConnect

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&connect)

	if err != nil {
		makeResponse(w, "failure")
		return
	}

	err = connect.Open()

	if err != nil {
		makeResponse(w, "connection refused")
		return
	}

	tables, err := connect.GetTables(parentSchema)

	if err != nil {
		makeResponse(w, "failure")
		return
	}

	defer connect.Close()

	response := "{\"status\":\"success\",\"tables\": ["

	for i, el := range tables {
		response += "\"" + el + "\""

		if i != len(tables)-1 {
			response += ","
		}
	}

	response += "]}"

	w.Write([]byte(response))
}

func SchemasList(w http.ResponseWriter, r *http.Request) {
	setSuccessHeader(w)

	var connect DBConnect

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&connect)

	if err != nil {
		makeResponse(w, "failure")
		return
	}

	err = connect.Open()

	if err != nil {
		makeResponse(w, "connection refused")
		return
	}

	defer connect.Close()

	schemas, err := connect.GetSchemas()

	if err != nil {
		makeResponse(w, "request failed")
	}

	response := "{\"status\":\"success\",\"schemas\": ["

	for i, el := range schemas {
		response += "\"" + el + "\""

		if i != len(schemas)-1 {
			response += ","
		}
	}

	response += "]}"

	w.Write([]byte(response))
}
