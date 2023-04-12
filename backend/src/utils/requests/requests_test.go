package requests

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	require.Equal(t, err, nil, "Error in request: %v", err)

	newRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	handler.ServeHTTP(newRecorder, req)
	require.Equal(t, http.StatusOK, newRecorder.Code, "handler returned wrong status code.\n Expected %v "+
		"actually %v", http.StatusOK, newRecorder.Code)
	require.Equal(t, "{\"status\":\"Success\"}\n", newRecorder.Body.String(), "handler returned "+
		"unexpected body. Expected %v actually %v", "{\"status\":\"Success\"}\n", newRecorder.Body.String())
	require.Equal(t, "application/json", newRecorder.Header().Get("Content-Type"))
}
