package repository

import (
	"errors"
	"sync"

	"github.com/Andressatass/trabalho-microinformatica-be/db/entities"
)

type MockUserRepository struct {
	data          map[uint]entities.UserInfo
	mu            sync.RWMutex
	autoIncrement uint
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		data:          make(map[uint]entities.UserInfo),
		autoIncrement: 1,
	}
}

func (r *MockUserRepository) Create(user entities.UserInfo) (uint, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.autoIncrement
	r.autoIncrement++

	r.data[user.ID] = user

	return user.ID, nil
}

func (r *MockUserRepository) FindById(id uint) (entities.UserInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.data[id]
	if !ok {
		return entities.UserInfo{}, errors.New("user not found")
	}

	return user, nil
}
