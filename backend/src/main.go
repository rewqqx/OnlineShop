package main

import (
	"backend/src/utils"
	"backend/src/utils/database"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Start Service:")

	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := "5432"
	if pgHost == "" {
		pgHost = "127.0.0.1"
		pgPort = "6000"
	}

	rdsHost := os.Getenv("REDIS_HOST")

	if rdsHost == "" {
		rdsHost = "127.0.0.1"
	}

	fmt.Println("Postgres Host: " + pgHost)
	fmt.Println("Redis Host: " + rdsHost)

	db := database.DBConnect{Ip: pgHost, Port: pgPort, Password: "pgpass", User: "postgres", Database: "postgres"}
	rds := database.New(rdsHost, "6379", "")

	err := db.Open()

	if err != nil {
		fmt.Println("<---- Can't Open Database ---->")
		panic(err)
	}

	err = rds.Ping()

	if err != nil {
		fmt.Println("<---- Can't Ping Redis ---->")
		panic(err)
	}

	fmt.Println("<---- Success Open Database ---->")

	server := utils.New(&db, rds)
	server.Start(9080)

	//prom.New()

}
