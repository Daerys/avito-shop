package usecase

import (
	"context"
	"fmt"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
)

type CoinUsecase interface {
	SendCoin(ctx context.Context, to, from string, amount int32) (*entity.CoinTransaction, error)
	RemoveCoins(ctx context.Context, from *entity.User, amount int) error
	GetHistory(ctx context.Context, user *entity.User) ([]entity.CoinTransaction, error)
}

type coinUsecaseImpl struct {
	repo repository.CoinRepository
}

func NewCoinUsecase(repo repository.CoinRepository) CoinUsecase {
	return &coinUsecaseImpl{
		repo: repo,
	}
}

func (u *coinUsecaseImpl) SendCoin(ctx context.Context, to, from string, amount int32) (*entity.CoinTransaction, error) {
	id, err := u.repo.SendCoin(ctx, to, from, amount)
	if err != nil {
		return nil, err
	}
	transaction, err := u.repo.GetTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (u *coinUsecaseImpl) RemoveCoins(ctx context.Context, from *entity.User, amount int) error {
	err := u.repo.RemoveCoins(ctx, from, amount)
	if err != nil {
		// Wrap the error with additional context for better debugging.
		return fmt.Errorf("failed to remove coins from user %s: %w", from, err)
	}
	return nil
}

func (u *coinUsecaseImpl) GetHistory(ctx context.Context, user *entity.User) ([]entity.CoinTransaction, error) {
	coinHistory, err := u.repo.GetHistory(ctx, user)
	if err != nil {
		// Wrap the error with additional context for better debugging.
		return nil, fmt.Errorf("failed to get history of user %s: %w", user.Username, err)
	}
	return coinHistory, nil
}
