package usecase_test

import (
	"sync"
	"testing"

	"avito-shop/internal/entity"
	"avito-shop/internal/repository/mocks"
	"avito-shop/internal/usecase/shop"
	"github.com/stretchr/testify/assert"
)

/*
TestShopUsecase_BuyItem verifies that a user can successfully purchase an item,
and that the user's coin balance and inventory are updated accordingly.
*/
func TestShopUsecase_BuyItem(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user := &entity.User{
		Username:  "testuser",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	err = shopUC.BuyItem("testuser", "t-shirt")
	assert.NoError(t, err)

	updatedUser, err := shopUC.GetUserInfo("testuser")
	assert.NoError(t, err)
	assert.Equal(t, 920, updatedUser.Coins)
	qty, ok := updatedUser.Inventory.Load("t-shirt")
	assert.True(t, ok)
	assert.Equal(t, 1, qty)
}

/*
TestShopUsecase_BuyItem_ItemNotFound verifies that attempting to purchase a non-existent item
returns an error indicating that the item was not found.
*/
func TestShopUsecase_BuyItem_ItemNotFound(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	// Создаем пользователя
	user := &entity.User{
		Username:  "testuser",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	err = shopUC.BuyItem("testuser", "nonexistent")
	assert.Error(t, err)
	assert.Equal(t, "item not found", err.Error())
}

/*
TestShopUsecase_BuyItem_InsufficientFunds verifies that purchasing an item with insufficient coins
results in an error.
*/
func TestShopUsecase_BuyItem_InsufficientFunds(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user := &entity.User{
		Username:  "testuser",
		Coins:     50,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	err = shopUC.BuyItem("testuser", "t-shirt")
	assert.Error(t, err)
	assert.Equal(t, "not enough coins to buy item", err.Error())
}

/*
TestShopUsecase_BuyMultipleItems verifies that a user can purchase the same item multiple times,
and that the cumulative cost and inventory quantity are correctly updated.
*/
func TestShopUsecase_BuyMultipleItems(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user := &entity.User{
		Username:  "testuser",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	err = shopUC.BuyItem("testuser", "t-shirt")
	assert.NoError(t, err)
	err = shopUC.BuyItem("testuser", "t-shirt")
	assert.NoError(t, err)

	updatedUser, err := shopUC.GetUserInfo("testuser")
	assert.NoError(t, err)
	assert.Equal(t, 840, updatedUser.Coins)
	qty, ok := updatedUser.Inventory.Load("t-shirt")
	assert.True(t, ok)
	assert.Equal(t, 2, qty)
}

/*
TestShopUsecase_SendCoins verifies that coins are successfully transferred between two users.
*/
func TestShopUsecase_SendCoins(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user1 := &entity.User{
		Username:  "user1",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	user2 := &entity.User{
		Username:  "user2",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user1)
	assert.NoError(t, err)
	_, err = userRepo.Create(user2)
	assert.NoError(t, err)

	err = shopUC.SendCoins("user1", "user2", 200)
	assert.NoError(t, err)

	updatedUser1, _ := shopUC.GetUserInfo("user1")
	updatedUser2, _ := shopUC.GetUserInfo("user2")
	assert.Equal(t, 800, updatedUser1.Coins)
	assert.Equal(t, 1200, updatedUser2.Coins)
}

/*
TestShopUsecase_SendCoins_InsufficientFunds verifies that a coin transfer fails if the sender does not have enough coins.
*/
func TestShopUsecase_SendCoins_InsufficientFunds(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user1 := &entity.User{
		Username:  "user1",
		Coins:     50,
		Inventory: &sync.Map{},
	}
	user2 := &entity.User{
		Username:  "user2",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user1)
	assert.NoError(t, err)
	_, err = userRepo.Create(user2)
	assert.NoError(t, err)

	err = shopUC.SendCoins("user1", "user2", 100)
	assert.Error(t, err)
	assert.Equal(t, "not enough coins", err.Error())
}

/*
TestShopUsecase_SendCoins_InvalidRecipient verifies that attempting to send coins to a non-existent recipient returns an error.
*/
func TestShopUsecase_SendCoins_InvalidRecipient(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user := &entity.User{
		Username:  "user1",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	err = shopUC.SendCoins("user1", "user2", 100)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

/*
TestShopUsecase_SendCoins_NegativeAmount verifies that attempting to send a negative coin amount returns an error.
*/
func TestShopUsecase_SendCoins_NegativeAmount(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user1 := &entity.User{
		Username:  "user1",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	user2 := &entity.User{
		Username:  "user2",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user1)
	assert.NoError(t, err)
	_, err = userRepo.Create(user2)
	assert.NoError(t, err)

	err = shopUC.SendCoins("user1", "user2", -50)
	assert.Error(t, err)
	assert.Equal(t, "amount must be positive", err.Error())
}

/*
TestShopUsecase_TransactionRecorded verifies that a coin transfer results in a recorded transaction.
*/
func TestShopUsecase_TransactionRecorded(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user1 := &entity.User{
		Username:  "user1",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	user2 := &entity.User{
		Username:  "user2",
		Coins:     1000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user1)
	assert.NoError(t, err)
	_, err = userRepo.Create(user2)
	assert.NoError(t, err)

	err = shopUC.SendCoins("user1", "user2", 150)
	assert.NoError(t, err)

	txs, err := txRepo.GetByUser("user1")
	assert.NoError(t, err)
	assert.Len(t, txs, 1)
	assert.Equal(t, "user1", txs[0].FromUser)
	assert.Equal(t, "user2", txs[0].ToUser)
	assert.Equal(t, 150, txs[0].Amount)
}

/*
TestConcurrentSendCoins verifies that concurrent coin transfers between two users result in the expected final balances.
*/
func TestConcurrentSendCoins(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user1 := &entity.User{
		Username:  "user1",
		Coins:     10000,
		Inventory: &sync.Map{},
	}
	user2 := &entity.User{
		Username:  "user2",
		Coins:     10000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user1)
	assert.NoError(t, err)
	_, err = userRepo.Create(user2)
	assert.NoError(t, err)

	const numTransfers = 100
	const transferAmount = 50

	var wg sync.WaitGroup
	wg.Add(numTransfers)

	for i := 0; i < numTransfers; i++ {
		go func() {
			defer wg.Done()
			err := shopUC.SendCoins("user1", "user2", transferAmount)
			assert.NoError(t, err)
		}()
	}
	wg.Wait()

	updatedUser1, err := shopUC.GetUserInfo("user1")
	assert.NoError(t, err)
	updatedUser2, err := shopUC.GetUserInfo("user2")
	assert.NoError(t, err)

	expectedUser1Balance := 10000 - numTransfers*transferAmount
	expectedUser2Balance := 10000 + numTransfers*transferAmount

	assert.Equal(t, expectedUser1Balance, updatedUser1.Coins)
	assert.Equal(t, expectedUser2Balance, updatedUser2.Coins)
}

/*
TestConcurrentBuyItem verifies that concurrent purchases of an item result in the expected final coin balance
and the correct number of items in the user's inventory.
*/
func TestConcurrentBuyItem(t *testing.T) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)

	user := &entity.User{
		Username:  "buyer",
		Coins:     10000,
		Inventory: &sync.Map{},
	}
	_, err := userRepo.Create(user)
	assert.NoError(t, err)

	const numPurchases = 50

	var wg sync.WaitGroup
	wg.Add(numPurchases)

	for i := 0; i < numPurchases; i++ {
		go func() {
			defer wg.Done()
			err := shopUC.BuyItem("buyer", "t-shirt")
			assert.NoError(t, err)
		}()
	}
	wg.Wait()

	updatedUser, err := shopUC.GetUserInfo("buyer")
	assert.NoError(t, err)

	expectedCoins := 10000 - numPurchases*80
	assert.Equal(t, expectedCoins, updatedUser.Coins)

	qty, exists := updatedUser.Inventory.Load("t-shirt")
	assert.True(t, exists)
	assert.Equal(t, numPurchases, qty)
}
