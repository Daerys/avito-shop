package postgres

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	var inventoryStr sql.NullString
	query := `SELECT id, username, password_hash, coins, inventory FROM users WHERE username=$1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Coins, &inventoryStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, err
	}
	user.Inventory = parseInventory(inventoryStr.String)
	return &user, nil
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {
	query := `INSERT INTO users (username, password_hash, coins, inventory)
				VALUES ($1, $2, $3, $4)
				ON CONFLICT (username) DO NOTHING
				RETURNING id;`
	invJSON, err := marshalInventory(user.Inventory)
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRow(query, user.Username, user.PasswordHash, user.Coins, invJSON).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	query := `UPDATE users SET coins=$1, inventory=$2 WHERE id=$3`
	invJSON, err := marshalInventory(user.Inventory)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, user.Coins, invJSON, user.ID)
	return err
}

func marshalInventory(inv *sync.Map) (string, error) {
	m := make(map[string]int)
	if inv != nil {
		inv.Range(func(key, value interface{}) bool {
			k, ok := key.(string)
			if !ok {
				return true
			}
			v, ok := value.(int)
			if !ok {
				return true
			}
			m[k] = v
			return true
		})
	}
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func parseInventory(data string) *sync.Map {
	var m map[string]int
	if data == "" {
		m = make(map[string]int)
	} else {
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			m = make(map[string]int)
		}
	}
	sm := &sync.Map{}
	for k, v := range m {
		sm.Store(k, v)
	}
	return sm
}
