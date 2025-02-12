package usecase

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"errors"
	"sync"
	"time"
)

// ShopUsecase encapsulates the business logic for the Avito Shop.
type ShopUsecase struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
	itemRepo        repository.ItemRepository
}

/*
NewShopUsecase creates and returns a new instance of ShopUsecase with the provided repositories.
*/
func NewShopUsecase(userRepo repository.UserRepository, transactionRepo repository.TransactionRepository, itemRepo repository.ItemRepository) *ShopUsecase {
	return &ShopUsecase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		itemRepo:        itemRepo,
	}
}

/*
GetUserInfo retrieves user information by username.
*/
func (s *ShopUsecase) GetUserInfo(username string) (*entity.User, error) {
	return s.userRepo.GetByUsername(username)
}

/*
CreateUser creates a new user in the system.
*/
func (s *ShopUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	return s.userRepo.Create(user)
}

/*
SendCoins transfers coins from one user to another.
It verifies that the amount is positive and that the sender has enough coins,
updates both users' coin balances, and records the transaction.
*/
func (s *ShopUsecase) SendCoins(fromUser, toUser string, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	sender, err := s.userRepo.GetByUsername(fromUser)
	if err != nil {
		return err
	}
	receiver, err := s.userRepo.GetByUsername(toUser)
	if err != nil {
		return err
	}
	if sender.Coins < amount {
		return errors.New("not enough coins")
	}
	sender.Coins -= amount
	receiver.Coins += amount

	if err := s.userRepo.Update(sender); err != nil {
		return err
	}
	if err := s.userRepo.Update(receiver); err != nil {
		return err
	}

	tx := entity.Transaction{
		FromUser:  fromUser,
		ToUser:    toUser,
		Amount:    amount,
		Type:      entity.TransactionTypeSend,
		Timestamp: time.Now(),
	}
	return s.transactionRepo.Create(tx)
}

/*
BuyItem processes the purchase of an item by a user.
It checks that the user has enough coins to buy the item, deducts the item price,
updates the user's inventory using a concurrent map, and records the transaction.
*/
func (s *ShopUsecase) BuyItem(username, itemName string) error {
	item, err := s.itemRepo.GetItem(itemName)
	if err != nil {
		return err
	}
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return err
	}
	if user.Coins < item.Price {
		return errors.New("not enough coins to buy item")
	}
	user.Coins -= item.Price

	if user.Inventory == nil {
		user.Inventory = &sync.Map{}
	}

	val, ok := user.Inventory.Load(item.Name)
	if !ok {
		user.Inventory.Store(item.Name, 1)
	} else {
		count := val.(int)
		user.Inventory.Store(item.Name, count+1)
	}

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	tx := entity.Transaction{
		FromUser:  username,
		ToUser:    "shop",
		Amount:    item.Price,
		Type:      entity.TransactionTypeBuy,
		Timestamp: time.Now(),
	}
	return s.transactionRepo.Create(tx)
}

/*
GetTransactions retrieves the transaction history for a given user.
*/
func (s *ShopUsecase) GetTransactions(username string) ([]entity.Transaction, error) {
	return s.transactionRepo.GetByUser(username)
}
