package mocks

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"errors"
	"sync"
)

/*
MockUserRepo implements the UserRepository interface using an in-memory map.
It uses a sync.RWMutex to ensure concurrent safety.
*/
type MockUserRepo struct {
	mu    sync.RWMutex
	Users map[string]*entity.User
}

/*
NewMockUserRepo creates and returns a new instance of MockUserRepo.
*/
func NewMockUserRepo() repository.UserRepository {
	return &MockUserRepo{
		Users: make(map[string]*entity.User),
	}
}

/*
GetByUsername retrieves a user by their username in a concurrent-safe manner.
Returns an error if the user is not found.
*/
func (f *MockUserRepo) GetByUsername(username string) (*entity.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	user, ok := f.Users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

/*
Create adds a new user to the repository in a concurrent-safe manner.
Returns an error if the user already exists.
*/
func (f *MockUserRepo) Create(user *entity.User) (*entity.User, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, exists := f.Users[user.Username]; exists {
		return nil, errors.New("user already exists")
	}
	f.Users[user.Username] = user
	return user, nil
}

/*
Update modifies an existing user in the repository in a concurrent-safe manner.
*/
func (f *MockUserRepo) Update(user *entity.User) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Users[user.Username] = user
	return nil
}
