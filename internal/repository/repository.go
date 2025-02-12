package repository

import "avito-shop/internal/entity"

/*
UserRepository defines the interface for user data operations.
*/
type UserRepository interface {
	GetByUsername(username string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) error
}

/*
TransactionRepository defines the interface for transaction operations.
*/
type TransactionRepository interface {
	Create(tx entity.Transaction) error
	GetByUser(username string) ([]entity.Transaction, error)
}

/*
ItemRepository defines the interface for item data operations.
*/
type ItemRepository interface {
	GetItem(name string) (*entity.Item, error)
	GetAllItems() ([]entity.Item, error)
}
