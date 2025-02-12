package mocks

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
)

/*
MockTransactionRepo implements the TransactionRepository interface using an in-memory slice.
*/
type MockTransactionRepo struct {
	Transactions []entity.Transaction
}

/*
NewMockTransactionRepo creates and returns a new instance of MockTransactionRepo.
*/
func NewMockTransactionRepo() repository.TransactionRepository {
	return &MockTransactionRepo{Transactions: []entity.Transaction{}}
}

/*
Create appends a new transaction to the in-memory transactions slice.
*/
func (f *MockTransactionRepo) Create(tx entity.Transaction) error {
	f.Transactions = append(f.Transactions, tx)
	return nil
}

/*
GetByUser retrieves all transactions associated with the given username.
It returns transactions where the user is either the sender or the recipient.
*/
func (f *MockTransactionRepo) GetByUser(username string) ([]entity.Transaction, error) {
	var result []entity.Transaction
	for _, tx := range f.Transactions {
		if tx.FromUser == username || tx.ToUser == username {
			result = append(result, tx)
		}
	}
	return result, nil
}
