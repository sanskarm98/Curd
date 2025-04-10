package test

import (
	"Curd/model"
	"errors"
)

// MockUserStore is an in-memory mock implementation of a user store.
// It uses a map to store users and an integer to track the next available user ID.
type MockUserStore struct {
	Users  map[int]model.User // Stores users with their ID as the key.
	NextID int                // Tracks the next available user ID.
}

// NewMockUserStore initializes and returns a new instance of MockUserStore.
func NewMockUserStore() *MockUserStore {
	return &MockUserStore{
		Users:  make(map[int]model.User), // Initialize the user map.
		NextID: 1,                        // Start IDs from 1.
	}
}

// CreateUser adds a new user to the store and assigns a unique ID to the user.
func (m *MockUserStore) CreateUser(user model.User) (model.User, error) {
	user.ID = m.NextID      // Assign the next available ID to the user.
	m.Users[user.ID] = user // Add the user to the map.
	m.NextID++              // Increment the next available ID.
	return user, nil        // Return the created user and no error.
}

// GetUser retrieves a user by their ID from the store.
// Returns an error if the user is not found.
func (m *MockUserStore) GetUser(id int) (model.User, error) {
	user, ok := m.Users[id] // Check if the user exists in the map.
	if !ok {
		return model.User{}, errors.New("user not found") // Return an error if not found.
	}
	return user, nil // Return the user and no error.
}

// GetAllUser retrieves all users from the store.
// Currently returns an empty slice (to be implemented further if needed).
func (m *MockUserStore) GetAllUser() ([]model.User, error) {
	return []model.User{}, nil // Return an empty slice and no error.
}

// UpdateUser updates an existing user's details in the store.
// Returns an error if the user is not found.
func (m *MockUserStore) UpdateUser(id int, user model.User) (model.User, error) {
	_, ok := m.Users[id] // Check if the user exists in the map.
	if !ok {
		return model.User{}, errors.New("user not found") // Return an error if not found.
	}
	user.ID = id       // Ensure the ID remains unchanged.
	m.Users[id] = user // Update the user in the map.
	return user, nil   // Return the updated user and no error.
}

// DeleteUser removes a user from the store by their ID.
// Returns an error if the user is not found.
func (m *MockUserStore) DeleteUser(id int) error {
	if _, ok := m.Users[id]; !ok { // Check if the user exists in the map.
		return errors.New("user not found") // Return an error if not found.
	}
	delete(m.Users, id) // Remove the user from the map.
	return nil          // Return no error.
}
