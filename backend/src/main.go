package main

import (
	"backend/src/utils"
	"fmt"
	"net/http"
	"os"
)

func main() {

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	fmt.Println("Postgres Host: " + host)

	database := utils.DBConnect{Ip: host, Port: "5432", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := database.Open()

	if err != nil {
		panic(err)
	}

	utils.SetDatabase(&database)

	userHandler := http.HandlerFunc(utils.GetUser)
	http.Handle("/users/", userHandler)

	pingHandler := http.HandlerFunc(utils.Ping)
	http.Handle("/", pingHandler)

	http.ListenAndServe(":8080", nil)
}
