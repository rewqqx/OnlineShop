package main

import (
	"backend/src/utils"
	"backend/src/utils/database"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Start Service:")

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	fmt.Println("Postgres Host: " + host)

	database := database.DBConnect{Ip: host, Port: "5432", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := database.Open()

	if err != nil {
		fmt.Println("<---- Can't Open Database ---->")
		panic(err)
	}

	fmt.Println("<---- Success Open Database ---->")

	server := utils.New(&database)
	server.Start(8080)

}
