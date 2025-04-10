package store

import (
	"Curd/model"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresUserStoreInterface defines the methods for interacting with the user store.
type PostgresUserStoreInterface interface {
	CreateUser(user model.User) (model.User, error)         // Create a new user in the database.
	GetUser(id int) (model.User, error)                     // Retrieve a user by their ID.
	UpdateUser(id int, user model.User) (model.User, error) // Update an existing user's details.
	DeleteUser(id int) error                                // Delete a user by their ID.
	GetAllUser() ([]model.User, error)                      // Retrieve all users from the database.
}

// PostgresUserStore is the implementation of PostgresUserStoreInterface using GORM.
type PostgresUserStore struct {
	db *gorm.DB // GORM database connection.
}

// NewPostgresUserStore initializes a new PostgresUserStore with the given DSN (Data Source Name).
func NewPostgresUserStore(dsn string) (*PostgresUserStore, error) {
	// Open a connection to the PostgreSQL database using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // Return an error if the connection fails.
	}

	// Auto-migrate the User schema to ensure the database structure matches the model.
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err // Return an error if migration fails.
	}

	// Return the initialized PostgresUserStore.
	return &PostgresUserStore{db: db}, nil
}

// CreateUser creates a new user in the database.
func (s *PostgresUserStore) CreateUser(user model.User) (model.User, error) {
	// Use GORM to insert the user into the database.
	if err := s.db.Create(&user).Error; err != nil {
		return model.User{}, err // Return an error if the operation fails.
	}
	return user, nil // Return the created user.
}

// GetUser retrieves a user by their ID.
func (s *PostgresUserStore) GetUser(id int) (model.User, error) {
	var user model.User
	// Use GORM to find the user by ID.
	if err := s.db.First(&user, id).Error; err != nil {
		return model.User{}, errors.New("user not found") // Return an error if the user is not found.
	}
	return user, nil // Return the retrieved user.
}

// GetAllUser retrieves all users from the database.
func (s *PostgresUserStore) GetAllUser() ([]model.User, error) {
	var users []model.User
	// Use GORM to retrieve all users.
	if err := s.db.Find(&users).Error; err != nil {
		return []model.User{}, errors.New("user not found") // Return an error if the operation fails.
	}
	return users, nil // Return the list of users.
}

// UpdateUser updates an existing user's details.
func (s *PostgresUserStore) UpdateUser(id int, updatedUser model.User) (model.User, error) {
	var user model.User
	// Use GORM to find the user by ID.
	if err := s.db.First(&user, id).Error; err != nil {
		return model.User{}, errors.New("user not found") // Return an error if the user is not found.
	}

	// Update the user's fields with the new data.
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	// Save the updated user back to the database.
	if err := s.db.Save(&user).Error; err != nil {
		return model.User{}, err // Return an error if the operation fails.
	}
	return user, nil // Return the updated user.
}

// DeleteUser deletes a user by their ID.
func (s *PostgresUserStore) DeleteUser(id int) error {
	// Use GORM to delete the user by ID.
	if err := s.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("user not found") // Return an error if the user is not found.
	}
	return nil // Return nil if the operation is successful.
}
