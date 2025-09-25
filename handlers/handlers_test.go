package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()

    Health(w, req)
    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("health status code = %d; want %d", resp.StatusCode, http.StatusOK)
    }
}

func TestAddHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/add?a=4&b=5", nil)
    w := httptest.NewRecorder()

    AddHandler(w, req)
    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("add status code = %d; want %d", resp.StatusCode, http.StatusOK)
    }
}
