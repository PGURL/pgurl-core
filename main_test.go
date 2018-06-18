package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, "pong", response["message"])
}
func TestShortRoute(t *testing.T) {
	router := SetupRouter()
	var jsonStr = []byte(`{"url": "https://www.google.com"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/short", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 6, len(response["shorturl"]))
	assert.Equal(t, "success", response["status"])
}
