package utils

import (
	"backend/src/utils/requests"
	"fmt"
	"net/http"
)

type Server struct {
}

func New() *Server {
	webServer := &Server{}
	webServer.prepare()
	return webServer
}

func (server *Server) prepare() {
	// Handle Image
	imageHandler := http.HandlerFunc(requests.GetImageRequest)
	http.Handle("/image/", imageHandler)

	// Bind Ping
	pingHandler := http.HandlerFunc(requests.Ping)
	http.Handle("/", pingHandler)
}

func (server *Server) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
