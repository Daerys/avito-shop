package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, username, password_hash, coins FROM users WHERE username = $1`
	if err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Coins); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) RegisterOrAuthenticate(ctx context.Context, username string, password string) (*entity.User, error) {
	var user entity.User
	query := `
        INSERT INTO users (username, password_hash, coins)
        VALUES ($1, $2, $3)
        ON CONFLICT (username)
        DO UPDATE SET username = EXCLUDED.username
        RETURNING id, username, password_hash, coins
    `

	err := r.db.QueryRowContext(ctx, query, username, password, 1000).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Coins)

	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	return &user, nil
}
