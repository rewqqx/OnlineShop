package main

import (
	"backend/src/utils"
	"backend/src/utils/requests"
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Start Service:")

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	fmt.Println("Postgres Host: " + host)

	database := utils.DBConnect{Ip: host, Port: "5432", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := database.Open()

	if err != nil {
		fmt.Println("<---- Can't Open Database ---->")
		panic(err)
	}

	fmt.Println("<---- Success Open Database ---->")

	requests.SetDatabase(&database)

	// Bind Users

	userHandler := http.HandlerFunc(requests.GetUser)
	http.Handle("/users/", userHandler)

	authHandler := http.HandlerFunc(requests.GetToken)
	http.Handle("/auth", authHandler)

	createHandler := http.HandlerFunc(requests.CreateUser)
	http.Handle("/users/create", createHandler)

	// Bind Items

	itemHandler := http.HandlerFunc(requests.GetItem)
	http.Handle("/items/", itemHandler)

	// Bind Ping

	pingHandler := http.HandlerFunc(requests.Ping)
	http.Handle("/", pingHandler)

	http.ListenAndServe(":8080", nil)
}
