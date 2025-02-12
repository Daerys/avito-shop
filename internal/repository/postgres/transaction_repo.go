package postgres

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"database/sql"
)

/*
TransactionRepository implements the repository.TransactionRepository interface using a PostgreSQL database.
*/
type TransactionRepository struct {
	db *sql.DB
}

/*
NewTransactionRepository creates and returns a new instance of TransactionRepository.
*/
func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return &TransactionRepository{db: db}
}

/*
Create inserts a new transaction record into the database.
*/
func (r *TransactionRepository) Create(tx entity.Transaction) error {
	query := `INSERT INTO transactions (from_user, to_user, amount, type, timestamp) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, tx.FromUser, tx.ToUser, tx.Amount, string(tx.Type), tx.Timestamp)
	return err
}

/*
GetByUser retrieves all transactions associated with the given username,
either as sender or receiver, ordered by timestamp in descending order.
*/
func (r *TransactionRepository) GetByUser(username string) ([]entity.Transaction, error) {
	query := `SELECT id, from_user, to_user, amount, type, timestamp FROM transactions WHERE from_user=$1 OR to_user=$1 ORDER BY timestamp DESC`
	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var tx entity.Transaction
		var txType string
		err := rows.Scan(&tx.ID, &tx.FromUser, &tx.ToUser, &tx.Amount, &txType, &tx.Timestamp)
		if err != nil {
			return nil, err
		}
		tx.Type = entity.TransactionType(txType)
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
