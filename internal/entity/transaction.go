package entity

import "time"

/*
TransactionType defines the type of a transaction.
*/
type TransactionType string

const (
	TransactionTypeSend TransactionType = "send"
	TransactionTypeBuy  TransactionType = "buy"
)

/*
Transaction represents a coin transaction operation.
*/
type Transaction struct {
	ID        int
	FromUser  string
	ToUser    string
	Amount    int
	Type      TransactionType
	Timestamp time.Time
}
