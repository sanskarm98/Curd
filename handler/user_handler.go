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

type UserHandler struct {
	Store store.UserStoreInterface
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/users":
		h.CreateUser(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/users/"):
		h.GetUser(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/users"):
		h.GetAllUser(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/users/"):
		h.UpdateUser(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/users/"):
		h.DeleteUser(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *UserHandler) extractID(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return 0, http.ErrMissingFile
	}
	return strconv.Atoi(parts[2])
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateUser Request: %s %s\n", r.Method, r.URL.Path)
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created, err := h.Store.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Send FCM Notification
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "New User Created",
			Body:  "User " + created.Name + " has been successfully created.",
		},
		Topic: "user-updates",
	}
	_, err = firebase.FCMClient.Send(r.Context(), message)
	if err != nil {
		log.Printf("Failed to send FCM notification: %v", err)
	}

	// Send Email Notification
	err = notification.SendEmail(
		created.Email,
		"Welcome to Our Service",
		"Hello "+created.Name+", welcome to our platform!",
	)
	if err != nil {
		log.Printf("Failed to send email notification: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetUser Request: %s %s\n", r.Method, r.URL.Path)
	id, err := h.extractID(r)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.Store.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllUser Request: %s %s\n", r.Method, r.URL.Path)
	users, err := h.Store.GetAllUser()
	if err != nil {
		http.Error(w, "no User found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateUser Request: %s %s\n", r.Method, r.URL.Path)
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updated, err := h.Store.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteUser Request: %s %s\n", r.Method, r.URL.Path)
	id, err := h.extractID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if err := h.Store.DeleteUser(id); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
