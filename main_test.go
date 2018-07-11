package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, "pong", response["message"])
}
func TestShortRoute(t *testing.T) {
	router := setupRouter()
	var response map[string]string
	var jsonStr = []byte(`{"url": "https://www.google.com"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/short", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 6, len(response["shorturl"]))
	assert.Equal(t, "success", response["status"])
}

func TestReShortRoute(t *testing.T) {
	baseURL := "https://www.google.com"
	router := setupRouter()
	var response map[string]string
	var jsonStr = []byte(fmt.Sprintf(`{"url": "%s"}`, baseURL))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/short", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	json.Unmarshal([]byte(w.Body.String()), &response)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/q/%s", response["shorturl"]), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, baseURL, w.HeaderMap["Location"][0])
	assert.Equal(t, 301, w.Code)
}
