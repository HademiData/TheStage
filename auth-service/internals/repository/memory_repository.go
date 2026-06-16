package repository

import (
	"auth-service/internals/domain"
	"errors"
)

type MemoryRepository struct {
	users map[string]domain.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[string]domain.User),
	}
}

func (m *MemoryRepository) Create(user domain.User) error {

	if _, ok := m.users[user.Email]; ok {
		return errors.New("User already exists")
	}

	m.users[user.Email] = user
	return nil
}

func (m *MemoryRepository) FindByEmail(email string) (*domain.User, error) {

	user, ok := m.users[email]

	if !ok {
		return nil, errors.New("User not found")
	}

	return &user, nil
}
