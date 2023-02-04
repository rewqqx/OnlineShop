package main

import (
	"backend/src/utils"
	"net/http"
)

func main() {
	schemasHandler := http.HandlerFunc(utils.SchemasList)
	http.Handle("/schemas", schemasHandler)

	tablesHandler := http.HandlerFunc(utils.TablesList)
	http.Handle("/tables/", tablesHandler)

	pingHandler := http.HandlerFunc(utils.Ping)
	http.Handle("/", pingHandler)

	http.ListenAndServe(":8080", nil)
}
