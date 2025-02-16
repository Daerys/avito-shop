package usecase

import (
	"context"
	"fmt"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
)

type ItemUsecase interface {
	GetItem(ctx context.Context, item string) (*entity.Item, error)
	AddItem(ctx context.Context, user *entity.User, item *entity.Item) error
	GetInventory(ctx context.Context, user *entity.User) ([]entity.InventoryItem, error)
}

type itemUsecaseImpl struct {
	repo repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) ItemUsecase {
	return &itemUsecaseImpl{
		repo: repo,
	}
}

func (i itemUsecaseImpl) GetItem(ctx context.Context, item string) (*entity.Item, error) {
	it := entity.Item{
		Name: item,
	}
	err := i.repo.GetItem(ctx, &it)
	if err != nil {
		return nil, fmt.Errorf("failed to find item %s: %w", item, err)
	}
	return &it, nil
}

func (i itemUsecaseImpl) AddItem(ctx context.Context, user *entity.User, item *entity.Item) error {
	err := i.repo.AddItem(ctx, user, item)
	if err != nil {
		return fmt.Errorf("failed to add item for %s: %w", user.Username, err)
	}
	return nil
}

func (i itemUsecaseImpl) GetInventory(ctx context.Context, user *entity.User) ([]entity.InventoryItem, error) {
	items, err := i.repo.GetInventory(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to find items for %s: %w", user.Username, err)
	}
	return items, nil
}
