package store

import (
	"errors"
	"hello_world1/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUserStoreInterface interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id int) (model.User, error)
	UpdateUser(id int, user model.User) (model.User, error)
	DeleteUser(id int) error
	GetAllUser() ([]model.User, error)
}
type PostgresUserStore struct {
	db *gorm.DB
}

func NewPostgresUserStore(dsn string) (*PostgresUserStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the User schema
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return &PostgresUserStore{db: db}, nil
}

func (s *PostgresUserStore) CreateUser(user model.User) (model.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *PostgresUserStore) GetUser(id int) (model.User, error) {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
func (s *PostgresUserStore) GetAllUser() ([]model.User, error) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		return []model.User{}, errors.New("user not found")
	}
	return users, nil
}
func (s *PostgresUserStore) UpdateUser(id int, updatedUser model.User) (model.User, error) {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return model.User{}, errors.New("user not found")
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := s.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *PostgresUserStore) DeleteUser(id int) error {
	if err := s.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}
