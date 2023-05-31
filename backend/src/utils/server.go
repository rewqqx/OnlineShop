package utils

import (
	"backend/src/utils/database"
	"backend/src/utils/prom"
	"backend/src/utils/requests"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Server struct {
	Database *database.DBConnect
	Redis    *database.Redis
}

func New(database *database.DBConnect, redis *database.Redis) *Server {
	webServer := &Server{Database: database, Redis: redis}
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

	// Bind Tags

	tagServer := requests.NewTagServer(server.Database)

	tagHandler := http.HandlerFunc(tagServer.GetTag)
	http.Handle("/tags/", tagHandler)

	// Bind Cart

	cartServer := requests.NewCartServer(server.Redis)

	cartHandler := http.HandlerFunc(cartServer.GetHandler)
	http.Handle("/cart/", cartHandler)

	// Bind Ping

	pingHandler := http.HandlerFunc(requests.Ping)
	http.Handle("/", pingHandler)

	// Bind Metrics
	//http.Handle("/metrics", promhttp.Handler())

	//ready
	prometheus.MustRegister(prom.MetricOnGETItems)
	prometheus.MustRegister(prom.MetricOnCreateItems)
	prometheus.MustRegister(prom.MetricOnPing)
	prometheus.MustRegister(prom.MetricOnGETTegs)
	prometheus.MustRegister(prom.MetricOnCreateUser)
	prometheus.MustRegister(prom.MetricOnGetUser)
	prometheus.MustRegister(prom.MetricOnUpdateUser)

	http.Handle("/metrics", promhttp.Handler())
	//go func() {
	//	http.ListenAndServe(":2112", nil)
	//
	//}()
}

func (server *Server) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

}
