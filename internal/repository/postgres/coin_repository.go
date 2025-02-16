package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
)

type coinRepositoryImpl struct {
	db *sql.DB
}

func NewCoinRepository(db *sql.DB) repository.CoinRepository {
	return &coinRepositoryImpl{db: db}
}

// RemoveCoins subtracts the specified amount of coins from the user's balance.
// It ensures that the balance does not fall below zero. If the balance is insufficient,
// an error is returned.
func (c *coinRepositoryImpl) RemoveCoins(ctx context.Context, from *entity.User, amount int) error {
	// Start a transaction to ensure atomicity.
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	// In case of error, rollback the transaction.
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Retrieve the current coin balance.
	var balance int
	err = tx.QueryRowContext(ctx, "SELECT coins FROM users WHERE id = $1", from.ID).Scan(&balance)
	if err != nil {
		return fmt.Errorf("failed to retrieve current balance for user %s: %w", from.Username, err)
	}

	// Check that the balance is sufficient.
	if balance < amount {
		return fmt.Errorf("insufficient balance: user %s has %d coins but %d required", from.Username, balance, amount)
	}

	// Subtract the coins from the user's balance.
	result, err := tx.ExecContext(ctx, "UPDATE users SET coins = coins - $1 WHERE id = $2", amount, from.ID)
	if err != nil {
		return fmt.Errorf("failed to update balance for user %s: %w", from.Username, err)
	}

	// Ensure that the update affected at least one row.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected for user %s: %w", from.Username, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated: user %s may not exist", from.Username)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	from.Coins = balance - amount

	return nil
}

func (c *coinRepositoryImpl) SendCoin(ctx context.Context, to, from string, amount int32) (int, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("ошибка начала транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var fromUserID int
	var fromCoins int32
	queryFrom := `SELECT id, coins FROM users WHERE username = $1 FOR UPDATE`
	if err = tx.QueryRowContext(ctx, queryFrom, from).Scan(&fromUserID, &fromCoins); err != nil {
		return -1, fmt.Errorf("ошибка получения данных отправителя: %w", err)
	}

	if fromCoins < amount {
		return -1, fmt.Errorf("недостаточно монет: имеется %d, требуется %d", fromCoins, amount)
	}

	var toUserID int
	queryTo := `SELECT id FROM users WHERE username = $1 FOR UPDATE`
	if err = tx.QueryRowContext(ctx, queryTo, to).Scan(&toUserID); err != nil {
		return -1, fmt.Errorf("ошибка получения данных получателя: %w", err)
	}

	updateFrom := `UPDATE users SET coins = coins - $1 WHERE id = $2`
	if _, err = tx.ExecContext(ctx, updateFrom, amount, fromUserID); err != nil {
		return -1, fmt.Errorf("ошибка обновления баланса отправителя: %w", err)
	}

	updateTo := `UPDATE users SET coins = coins + $1 WHERE id = $2`
	if _, err = tx.ExecContext(ctx, updateTo, amount, toUserID); err != nil {
		return -1, fmt.Errorf("ошибка обновления баланса получателя: %w", err)
	}
	var transactionID int
	insertTransaction := `INSERT INTO coin_transactions (from_user_id, to_user_id, amount) VALUES ($1, $2, $3) RETURNING id`
	if err = tx.QueryRowContext(ctx, insertTransaction, fromUserID, toUserID, amount).Scan(&transactionID); err != nil {
		return -1, fmt.Errorf("ошибка записи транзакции: %w", err)
	}

	// Фиксируем транзакцию
	if err = tx.Commit(); err != nil {
		return -1, fmt.Errorf("ошибка фиксации транзакции: %w", err)
	}

	return transactionID, nil
}

func (c *coinRepositoryImpl) GetTransaction(ctx context.Context, id int) (*entity.CoinTransaction, error) {
	var transaction entity.CoinTransaction
	query := `SELECT id, from_user_id, to_user_id, amount FROM coin_transactions WHERE id = $1 FOR UPDATE`
	if err := c.db.QueryRowContext(ctx, query, id).Scan(
		&transaction.ID,
		&transaction.FromUserID,
		&transaction.ToUserID,
		&transaction.Amount,
	); err != nil {
		return nil, fmt.Errorf("ошибка получения данных транзакции: %w", err)
	}
	return &transaction, nil
}

func (c *coinRepositoryImpl) GetHistory(ctx context.Context, user *entity.User) ([]entity.CoinTransaction, error) {
	const query = `
        SELECT 
            id, 
            from_user_id, 
            to_user_id, 
            amount 
        FROM coin_transactions 
        WHERE 
            from_user_id = $1 OR 
            to_user_id = $1 
        ORDER BY id DESC`

	var transactions []entity.CoinTransaction

	// Выполняем запрос с таймаутом из контекста
	rows, err := c.db.QueryContext(ctx, query, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.CoinTransaction{}, nil
		}
		return nil, fmt.Errorf("failed to get coin history: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var invItem entity.CoinTransaction
		if err = rows.Scan(&invItem.ID, &invItem.FromUserID, &invItem.ToUserID, &invItem.Amount); err != nil {
			return nil, fmt.Errorf("failed to scan coin history: %w", err)
		}
		transactions = append(transactions, invItem)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating coin history: %w", err)
	}
	return transactions, nil
}
