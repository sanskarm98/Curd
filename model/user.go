package model

// User represents the structure of a user in the system.
// It includes fields for ID, Name, and Email.
type User struct {
	// ID is the primary key for the User table in the database.
	// It is auto-incremented by GORM.
	ID int `gorm:"primaryKey;autoIncrement"`

	// Name is the name of the user.
	Name string `json:"name"`

	// Email is the email address of the user.
	Email string `json:"email"`
}
