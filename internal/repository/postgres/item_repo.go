package postgres

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"database/sql"
	"errors"
	"fmt"
)

/*
ItemRepository implements the repository.ItemRepository interface using a PostgreSQL database.
*/
type ItemRepository struct {
	db *sql.DB
}

/*
NewItemRepository creates and returns a new instance of ItemRepository.
NOTE: Although an in-memory repository might suffice if the list of items is fixed, using a database
provides more flexibility.
*/
func NewItemRepository(db *sql.DB) repository.ItemRepository {
	return &ItemRepository{db: db}
}

/*
GetItem retrieves an item by its name from the database.
It returns an error if the item is not found.
*/
func (r *ItemRepository) GetItem(name string) (*entity.Item, error) {
	var item entity.Item
	query := "SELECT name, price FROM items WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&item.Name, &item.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, fmt.Errorf("failed to get item: %w", err)
	}
	return &item, nil
}

/*
GetAllItems retrieves all items from the database.
It returns a slice of items or an error if the query fails.
*/
func (r *ItemRepository) GetAllItems() ([]entity.Item, error) {
	query := "SELECT name, price FROM items"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}
	defer rows.Close()

	var items []entity.Item
	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.Name, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return items, nil
}
