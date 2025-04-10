package store

import (
	"Curd/model"
	"errors"
	"log"
	"sync"
)

// UserStoreInterface defines the methods that a UserStore must implement.
type UserStoreInterface interface {
	CreateUser(user model.User) (model.User, error)         // Create a new user
	GetUser(id int) (model.User, error)                     // Retrieve a user by ID
	UpdateUser(id int, user model.User) (model.User, error) // Update an existing user by ID
	DeleteUser(id int) error                                // Delete a user by ID
	GetAllUser() ([]model.User, error)                      // Retrieve all users
}

// UserStore is an in-memory implementation of UserStoreInterface.
// It uses a map to store users and a mutex for thread-safe operations.
type UserStore struct {
	sync.Mutex
	users  map[int]model.User // Map to store users with their ID as the key
	nextID int                // Counter to generate unique user IDs
}

// NewUserStore initializes and returns a new UserStore instance.
func NewUserStore() (*UserStore, error) {
	return &UserStore{
		users:  make(map[int]model.User), // Initialize the user map
		nextID: 1,                        // Start IDs from 1
	}, nil
}

// CreateUser adds a new user to the store and assigns a unique ID.
func (s *UserStore) CreateUser(user model.User) (model.User, error) {
	s.Lock()
	defer s.Unlock()

	// Assign a unique ID to the user
	user.ID = s.nextID
	s.users[user.ID] = user // Add the user to the map
	s.nextID++              // Increment the ID counter
	return user, nil
}

// GetUser retrieves a user by their ID.
func (s *UserStore) GetUser(id int) (model.User, error) {
	s.Lock()
	defer s.Unlock()

	log.Printf("ID: %v\n", id) // Log the ID being retrieved
	user, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("user not found") // Return an error if the user doesn't exist
	}
	return user, nil
}

// GetAllUser retrieves all users from the store.
func (s *UserStore) GetAllUser() ([]model.User, error) {
	s.Lock()
	defer s.Unlock()

	var users []model.User
	for _, value := range s.users {
		users = append(users, value) // Collect all users into a slice
	}
	if len(users) == 0 {
		return nil, errors.New("no users found") // Return an error if no users exist
	}
	return users, nil
}

// UpdateUser updates an existing user's details by their ID.
func (s *UserStore) UpdateUser(id int, user model.User) (model.User, error) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("user not found") // Return an error if the user doesn't exist
	}
	user.ID = id       // Ensure the ID remains unchanged
	s.users[id] = user // Update the user in the map
	return user, nil
}

// DeleteUser removes a user from the store by their ID.
func (s *UserStore) DeleteUser(id int) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.users[id]; !ok {
		return errors.New("user not found") // Return an error if the user doesn't exist
	}
	delete(s.users, id) // Remove the user from the map
	return nil
}
