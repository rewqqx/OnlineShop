package utils

import (
	"backend/src/utils/database"
	"backend/src/utils/requests"
	"fmt"
	"net/http"
)

type Server struct {
	Database *database.DBConnect
}

func New(database *database.DBConnect) *Server {
	webServer := &Server{Database: database}
	webServer.prepare()
	return webServer
}

func (server *Server) prepare() {

	userServer := requests.NewUserServer(server.Database)

	userHandler := http.HandlerFunc(userServer.GetUser)
	http.Handle("/users/", userHandler)

	authHandler := http.HandlerFunc(userServer.GetToken)
	http.Handle("/auth", authHandler)

	createHandler := http.HandlerFunc(userServer.CreateUser)
	http.Handle("/users/create", createHandler)

	updateHandler := http.HandlerFunc(userServer.UpdateUser)
	http.Handle("/users/update/", updateHandler)

	// Bind Items

	itemServer := requests.NewItemServer(server.Database)

	itemHandler := http.HandlerFunc(itemServer.GetItem)
	http.Handle("/items/", itemHandler)

	// Bind Ping

	pingHandler := http.HandlerFunc(requests.Ping)
	http.Handle("/", pingHandler)
}

func (server *Server) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
