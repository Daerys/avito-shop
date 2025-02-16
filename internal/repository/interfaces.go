package repository

import (
	"context"
	"github.com/Daerys/avito-shop/internal/entity"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	RegisterOrAuthenticate(ctx context.Context, username string, hash string) (*entity.User, error)
}

type CoinRepository interface {
	SendCoin(ctx context.Context, to, from string, amount int32) (int, error)
	GetTransaction(ctx context.Context, id int) (*entity.CoinTransaction, error)
	RemoveCoins(ctx context.Context, from *entity.User, amount int) error
	GetHistory(ctx context.Context, user *entity.User) ([]entity.CoinTransaction, error)
}

type ItemRepository interface {
	GetItem(ctx context.Context, item *entity.Item) error
	AddItem(ctx context.Context, user *entity.User, item *entity.Item) error
	GetInventory(ctx context.Context, user *entity.User) ([]entity.InventoryItem, error)
}
