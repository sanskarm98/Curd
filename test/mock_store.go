package test

import (
	"Curd/model"
	"errors"
)

type MockUserStore struct {
	Users  map[int]model.User
	NextID int
}

func NewMockUserStore() *MockUserStore {
	return &MockUserStore{
		Users:  make(map[int]model.User),
		NextID: 1,
	}
}

func (m *MockUserStore) CreateUser(user model.User) (model.User, error) {
	user.ID = m.NextID
	m.Users[user.ID] = user
	m.NextID++
	return user, nil
}

func (m *MockUserStore) GetUser(id int) (model.User, error) {
	user, ok := m.Users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
func (m *MockUserStore) GetAllUser() ([]model.User, error) {

	return []model.User{}, nil
}
func (m *MockUserStore) UpdateUser(id int, user model.User) (model.User, error) {
	_, ok := m.Users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	user.ID = id
	m.Users[id] = user
	return user, nil
}

func (m *MockUserStore) DeleteUser(id int) error {
	if _, ok := m.Users[id]; !ok {
		return errors.New("user not found")
	}
	delete(m.Users, id)
	return nil
}
