package handler

import (
	"Curd/firebase"
	"Curd/model"
	"Curd/notification"
	"Curd/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"firebase.google.com/go/messaging"
)

// UserHandler handles HTTP requests related to user operations.
type UserHandler struct {
	Store store.UserStoreInterface // Interface for user data storage operations.
}

// ServeHTTP routes incoming HTTP requests to the appropriate handler function.
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Route requests based on HTTP method and URL path.
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/users":
		h.CreateUser(w, r) // Handle user creation.
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/users/"):
		h.GetUser(w, r) // Handle fetching a single user by ID.
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/users"):
		h.GetAllUser(w, r) // Handle fetching all users.
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/users/"):
		h.UpdateUser(w, r) // Handle updating a user by ID.
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/users/"):
		h.DeleteUser(w, r) // Handle deleting a user by ID.
	default:
		http.NotFound(w, r) // Return 404 for unsupported routes.
	}
}

// extractID extracts the user ID from the URL path.
func (h *UserHandler) extractID(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return 0, http.ErrMissingFile // Return error if ID is missing.
	}
	return strconv.Atoi(parts[2]) // Convert ID to integer.
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateUser Request: %s %s\n", r.Method, r.URL.Path)

	// Decode the request body into a User object.
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest) // Return 400 for invalid input.
		return
	}

	// Create the user in the data store.
	created, err := h.Store.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError) // Return 500 for server error.
		return
	}

	// Send FCM Notification for the new user.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "New User Created",
			Body:  "User " + created.Name + " has been successfully created.",
		},
		Topic: "user-updates",
	}
	_, err = firebase.FCMClient.Send(r.Context(), message)
	if err != nil {
		log.Printf("Failed to send FCM notification: %v", err) // Log FCM notification failure.
	}

	// Send Email Notification to the new user.
	err = notification.SendEmail(
		created.Email,
		"Welcome to Our Service",
		"Hello "+created.Name+", welcome to our platform!",
	)
	if err != nil {
		log.Printf("Failed to send email notification: %v", err) // Log email notification failure.
	}

	// Respond with the created user object and HTTP 201 status.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// GetUser handles fetching a single user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetUser Request: %s %s\n", r.Method, r.URL.Path)

	// Extract the user ID from the URL path.
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return 400 for invalid ID.
		return
	}

	// Fetch the user from the data store.
	user, err := h.Store.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound) // Return 404 if user is not found.
		return
	}

	// Respond with the user object.
	json.NewEncoder(w).Encode(user)
}

// GetAllUser handles fetching all users.
func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllUser Request: %s %s\n", r.Method, r.URL.Path)

	// Fetch all users from the data store.
	users, err := h.Store.GetAllUser()
	if err != nil {
		http.Error(w, "No users found", http.StatusNotFound) // Return 404 if no users are found.
		return
	}

	// Respond with the list of users.
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles updating a user by ID.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateUser Request: %s %s\n", r.Method, r.URL.Path)

	// Extract the user ID from the URL path.
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return 400 for invalid ID.
		return
	}

	// Decode the request body into a User object.
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest) // Return 400 for invalid input.
		return
	}

	// Update the user in the data store.
	updated, err := h.Store.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound) // Return 404 if user is not found.
		return
	}

	// Respond with the updated user object.
	json.NewEncoder(w).Encode(updated)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteUser Request: %s %s\n", r.Method, r.URL.Path)

	// Extract the user ID from the URL path.
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return 400 for invalid ID.
		return
	}

	// Delete the user from the data store.
	if err := h.Store.DeleteUser(id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound) // Return 404 if user is not found.
		return
	}

	// Respond with HTTP 204 No Content status.
	w.WriteHeader(http.StatusNoContent)
}
