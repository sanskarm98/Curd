package store

import (
	"errors"
	"hello_world1/model"
	"log"
	"sync"
)

type UserStoreInterface interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id int) (model.User, error)
	UpdateUser(id int, user model.User) (model.User, error)
	DeleteUser(id int) error
	GetAllUser() ([]model.User, error)
}
type UserStore struct {
	sync.Mutex
	users  map[int]model.User
	nextID int
}

func NewUserStore() (*UserStore, error) {
	return &UserStore{
		users:  make(map[int]model.User),
		nextID: 1,
	}, nil
}

func (s *UserStore) CreateUser(user model.User) (model.User, error) {
	s.Lock()
	defer s.Unlock()
	user.ID = s.nextID
	s.users[user.ID] = user
	s.nextID++
	return user, nil
}

func (s *UserStore) GetUser(id int) (model.User, error) {
	s.Lock()
	defer s.Unlock()
	log.Printf("ID: %v\n", id)
	user, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
func (s *UserStore) GetAllUser() ([]model.User, error) {
	s.Lock()
	defer s.Unlock()
	var users []model.User
	for _, value := range s.users {
		users = append(users, value)
	}
	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}
func (s *UserStore) UpdateUser(id int, user model.User) (model.User, error) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	user.ID = id
	s.users[id] = user
	return user, nil
}

func (s *UserStore) DeleteUser(id int) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
