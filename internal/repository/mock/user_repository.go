package mock

import (
	"context"
	"github.com/Daerys/avito-shop/internal/entity"
)

// UserRepository - заглушка для UserRepository
type UserRepository struct {
	users *map[string]*entity.User
}

func NewUserRepository(users *map[string]*entity.User) *UserRepository {
	return &UserRepository{
		users: users,
	}
}

func (m *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, exists := (*m.users)[username]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (m *UserRepository) RegisterOrAuthenticate(ctx context.Context, username string, password string) (*entity.User, error) {
	if user, exists := (*m.users)[username]; exists {
		return user, nil
	}
	newUser := &entity.User{Username: username, PasswordHash: password, Coins: 1000}
	(*m.users)[username] = newUser
	return newUser, nil
}
