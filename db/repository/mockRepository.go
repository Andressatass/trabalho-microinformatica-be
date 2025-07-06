package repository

import (
	"errors"
	"sync"

	"github.com/Andressatass/trabalho-microinformatica-be/trabalho-microinformatica-be/db/entities"
)

type MockUserRepository struct {
	data          map[string]entities.UserInfo
	mu            sync.RWMutex
	autoIncrement uint
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		data:          make(map[string]entities.UserInfo),
		autoIncrement: 1,
	}
}

func (r *MockUserRepository) Create(user entities.UserInfo) (uint, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.autoIncrement
	r.autoIncrement++

	r.data[user.UUID] = user

	return user.ID, nil
}

func (r *MockUserRepository) FindById(id string) (entities.UserInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.data[id]
	if !ok {
		return entities.UserInfo{}, errors.New("user not found")
	}

	return user, nil
}
