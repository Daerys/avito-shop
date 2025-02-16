package mock

import (
	"context"
	"errors"
	"github.com/Daerys/avito-shop/internal/entity"
	"strconv"
)

type CoinRepository struct {
	balances map[string]int
	users    *map[string]*entity.User
	history  []*entity.CoinTransaction
}

func NewCoinRepository(users *map[string]*entity.User) *CoinRepository {
	return &CoinRepository{
		balances: make(map[string]int),
		users:    users,
		history:  make([]*entity.CoinTransaction, 0),
	}
}

func (m *CoinRepository) RemoveCoins(ctx context.Context, from *entity.User, amount int) error {
	if from.Coins < amount {
		return errors.New("insufficient balance")
	}
	from.Coins -= amount
	(*m.users)[from.Username].Coins -= amount
	return nil
}

func (m *CoinRepository) SendCoin(ctx context.Context, to, from string, amount int32) (int, error) {
	if (*m.users)[from].Coins < int(amount) {
		return -1, errors.New("not enough coins")
	}
	(*m.users)[from].Coins -= int(amount)
	(*m.users)[to].Coins += int(amount)
	fromUserID, _ := strconv.ParseInt(from, 10, 64)
	toUserID, _ := strconv.ParseInt(to, 10, 64)
	m.history = append(m.history, &entity.CoinTransaction{
		ID:         int64(len(m.history)),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Amount:     int(amount),
	})
	return len(m.history) - 1, nil
}

func (m *CoinRepository) GetTransaction(ctx context.Context, id int) (*entity.CoinTransaction, error) {
	if id < 0 || id >= len(m.history) {
		return nil, errors.New("transaction not found")
	}
	return m.history[id], nil
}

func (m *CoinRepository) GetHistory(ctx context.Context, user *entity.User) ([]entity.CoinTransaction, error) {
	var result []entity.CoinTransaction
	for _, transaction := range m.history {
		if transaction.FromUserID == user.ID || transaction.ToUserID == user.ID {
			result = append(result, *transaction)
		}
	}
	return result, nil
}
