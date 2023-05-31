package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

// ready
var MetricOnGETItems = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_get_items",
		Help: "Количество запросов на получение предмета/предметов",
	})

var MetricOnCreateItems = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_create_item",
		Help: "Количество запросов на создание предмета",
	})

var MetricOnPing = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_ping",
		Help: "Количество запросов на Ping сервиса",
	})

var MetricOnGETTegs = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_get_tegs",
		Help: "Количество запросов на получение тэгов",
	})

var MetricOnCreateUser = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_create_user",
		Help: "Количество запросов на создание пользователя",
	})

var MetricOnGetUser = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_get_user",
		Help: "Количество запросов на поулчние пользователя",
	})

var MetricOnUpdateUser = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_request_update_user",
		Help: "Количество запросов на обновления данных пользователя",
	})

//func New() {
//	prometheus.MustRegister(MetricOnGET)
//
//	go func() {
//		http.Handle("/metrics", promhttp.Handler())
//		http.ListenAndServe(":2112", nil)
//
//	}()
//}

//var (
//	GetRequests = prometheus.NewCounter(
//		prometheus.CounterOpts{
//			Name: "get_requests",
//			Help: "Количество запросов на переход по shortUrl",
//		},
//	)
//)
//
//func init() {
//	prometheus.MustRegister(GetRequests)
//}
