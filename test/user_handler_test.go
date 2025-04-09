package test

import (
	"bytes"
	"encoding/json"
	"hello_world1/handler"
	"hello_world1/model"
	"hello_world1/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMockServer() http.Handler {
	mockStore := NewMockUserStore()
	handler := &handler.UserHandler{Store: mockStore}
	return router.NewRouter(handler)
}

func TestCreateUser(t *testing.T) {
	server := setupMockServer()

	user := model.User{Name: "Alice", Email: "alice@example.com"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", w.Code)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	server := setupMockServer()

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", w.Code)
	}
}
