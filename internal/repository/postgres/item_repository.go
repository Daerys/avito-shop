package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
)

type itemRepositoryImpl struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) repository.ItemRepository {
	return &itemRepositoryImpl{db: db}
}

func (i *itemRepositoryImpl) GetItem(ctx context.Context, item *entity.Item) error {
	query := `SELECT id, name, price FROM items WHERE name = $1 FOR UPDATE`
	if err := i.db.QueryRowContext(ctx, query, item.Name).Scan(
		&item.ID,
		&item.Name,
		&item.Price,
	); err != nil {
		return fmt.Errorf("couldn't find item %s: %w", item.Name, err)
	}
	return nil
}

// AddItem adds an item to the user's inventory. It updates the inventory table by increasing
// the quantity for the given item if it exists, or inserts a new record if it does not.
// It then retrieves all inventory items for the user and assigns them to the items pointer.
func (i *itemRepositoryImpl) AddItem(ctx context.Context, user *entity.User, item *entity.Item) error {
	// Begin a transaction to ensure atomicity.
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	// Rollback the transaction if any error occurs.
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Attempt to update the inventory record if the item already exists for the user.
	res, err := tx.ExecContext(ctx, `
		UPDATE inventory 
		SET quantity = quantity + 1 
		WHERE user_id = $1 AND item_type = $2
	`, user.ID, item.Name)
	if err != nil {
		return fmt.Errorf("failed to update inventory item: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	// If no record was updated, the item doesn't exist yet; insert a new inventory record.
	if rowsAffected == 0 {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO inventory (user_id, item_type, quantity) 
			VALUES ($1, $2, 1)
		`, user.ID, item.Name)
		if err != nil {
			return fmt.Errorf("failed to insert new inventory item: %w", err)
		}
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (i *itemRepositoryImpl) GetInventory(ctx context.Context, user *entity.User) ([]entity.InventoryItem, error) {
	// Retrieve all inventory items for the user.
	rows, err := i.db.QueryContext(ctx, `
		SELECT user_id, item_type, quantity 
		FROM inventory 
		WHERE user_id = $1
	`, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.InventoryItem{}, nil
		}
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}
	defer rows.Close()

	var inventoryItems []entity.InventoryItem
	for rows.Next() {
		var invItem entity.InventoryItem
		if err = rows.Scan(&invItem.UserID, &invItem.ItemType, &invItem.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan inventory item: %w", err)
		}
		inventoryItems = append(inventoryItems, invItem)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating inventory items: %w", err)
	}
	return inventoryItems, nil
}
