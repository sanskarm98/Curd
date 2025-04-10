package test

import (
	"Curd/handler"
	"Curd/model"
	"Curd/router"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setupMockServer initializes a mock server with a mock user store and returns the router.
func setupMockServer() http.Handler {
	mockStore := NewMockUserStore()                   // Mock implementation of the user store.
	handler := &handler.UserHandler{Store: mockStore} // UserHandler with the mock store.
	return router.NewRouter(handler)                  // Return a new router with the handler.
}

// TestCreateUser tests the creation of a new user.
func TestCreateUser(t *testing.T) {
	server := setupMockServer() // Set up the mock server.

	// Create a sample user to send in the request body.
	user := model.User{Name: "Alice", Email: "alice@example.com"}
	body, _ := json.Marshal(user) // Serialize the user object to JSON.

	// Create a POST request to the /users endpoint with the user data.
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	w := httptest.NewRecorder() // Recorder to capture the response.

	server.ServeHTTP(w, req) // Serve the request.

	// Assert that the response status code is 201 Created.
	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", w.Code)
	}
}

// TestGetUserError tests the scenario where a user is not found (404 Not Found).
func TestGetUserError(t *testing.T) {
	server := setupMockServer() // Set up the mock server.

	// Create a GET request to fetch a user that doesn't exist.
	req := httptest.NewRequest(http.MethodGet, "/users/1", bytes.NewReader(nil))
	w := httptest.NewRecorder() // Recorder to capture the response.

	server.ServeHTTP(w, req) // Serve the request.

	// Assert that the response status code is 404 Not Found.
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 Got the user, got %d", w.Code)
	}
}

// TestGetUser_NotFound tests the scenario where a specific user ID is not found.
func TestGetUser_NotFound(t *testing.T) {
	server := setupMockServer() // Set up the mock server.

	// Create a GET request to fetch a user with an ID that doesn't exist.
	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder() // Recorder to capture the response.

	server.ServeHTTP(w, req) // Serve the request.

	// Assert that the response status code is 404 Not Found.
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", w.Code)
	}
}
