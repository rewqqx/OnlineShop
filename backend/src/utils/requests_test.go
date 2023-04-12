package utils

import (
	"backend/src/utils/database"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	host := "127.0.0.1"
	database := database.DBConnect{Ip: host, Port: "5432", Password: "pgpass", User: "postgres", Database: "postgres"}
	err := database.Open()

	require.Equal(t, nil, err)

	server := New(&database)
	go server.Start(8080)

	resp, err := http.Get("http://" + host + ":8080/")

	require.Equal(t, err, nil, "Error in request: %v", err)
	require.Equal(t, resp.StatusCode, http.StatusOK)
}
