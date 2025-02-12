package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	httpDelivery "avito-shop/internal/delivery/http"
	"avito-shop/internal/repository/mocks"
	"avito-shop/internal/usecase/shop"
)

// setupRouter initializes the fake repositories, business logic (use case),
// and registers all HTTP routes using Gin. It returns the router, the use case,
// and the JWT secret used.
func setupRouter() (*gin.Engine, *usecase.ShopUsecase, string) {
	userRepo := mocks.NewMockUserRepo()
	txRepo := mocks.NewMockTransactionRepo()
	itemRepo := mocks.NewMockItemRepo()
	shopUC := usecase.NewShopUsecase(userRepo, txRepo, itemRepo)
	router := gin.Default()
	jwtSecret := "testsecret"
	httpDelivery.RegisterRoutes(router, shopUC, jwtSecret)
	return router, shopUC, jwtSecret
}

/*
TestIntegration_AuthAndGetInfo tests the authentication endpoint (/api/auth)
to create a new user and then calls /api/info to verify that the new user has
an initial balance of 1000 coins and an empty inventory.
*/
func TestIntegration_AuthAndGetInfo(t *testing.T) {
	router, _, _ := setupRouter()

	authReqBody := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	authJSON, _ := json.Marshal(authReqBody)
	req, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var authResp map[string]string
	json.Unmarshal(rec.Body.Bytes(), &authResp)
	token, ok := authResp["token"]
	assert.True(t, ok)
	assert.NotEmpty(t, token)

	reqInfo, _ := http.NewRequest("GET", "/api/info", nil)
	reqInfo.Header.Set("Authorization", "Bearer "+token)
	recInfo := httptest.NewRecorder()
	router.ServeHTTP(recInfo, reqInfo)
	assert.Equal(t, http.StatusOK, recInfo.Code)

	var infoResp map[string]interface{}
	json.Unmarshal(recInfo.Body.Bytes(), &infoResp)
	assert.Equal(t, float64(1000), infoResp["coins"])
}

/*
TestIntegration_BuyItem tests that a user can buy an item ("/api/buy/{item}")
by authenticating via /api/auth, buying an item (e.g., "t-shirt"),
and then verifying that the user's coin balance is reduced and the inventory is updated.
*/
func TestIntegration_BuyItem(t *testing.T) {
	router, shopUC, _ := setupRouter()

	authReqBody := map[string]string{"username": "buyer", "password": "buytest"}
	authJSON, _ := json.Marshal(authReqBody)
	reqAuth, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON))
	reqAuth.Header.Set("Content-Type", "application/json")
	recAuth := httptest.NewRecorder()
	router.ServeHTTP(recAuth, reqAuth)
	assert.Equal(t, http.StatusOK, recAuth.Code)
	var authResp map[string]string
	json.Unmarshal(recAuth.Body.Bytes(), &authResp)
	token := authResp["token"]
	assert.NotEmpty(t, token)

	reqBuy, _ := http.NewRequest("GET", "/api/buy/t-shirt", nil)
	reqBuy.Header.Set("Authorization", "Bearer "+token)
	recBuy := httptest.NewRecorder()
	router.ServeHTTP(recBuy, reqBuy)
	assert.Equal(t, http.StatusOK, recBuy.Code)

	user, err := shopUC.GetUserInfo("buyer")
	assert.NoError(t, err)
	assert.Equal(t, 920, user.Coins)
	qty, exists := user.Inventory.Load("t-shirt")
	assert.True(t, exists)
	assert.Equal(t, 1, qty)
}

/*
TestIntegration_SendCoin tests the coin transfer endpoint (/api/sendCoin) by
authenticating two users, performing a coin transfer from user1 to user2,
and then verifying that their balances are updated accordingly.
*/
func TestIntegration_SendCoin(t *testing.T) {
	router, shopUC, _ := setupRouter()

	authReq1 := map[string]string{"username": "user1", "password": "pass1"}
	authJSON1, _ := json.Marshal(authReq1)
	req1, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON1))
	req1.Header.Set("Content-Type", "application/json")
	rec1 := httptest.NewRecorder()
	router.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	var authResp1 map[string]string
	json.Unmarshal(rec1.Body.Bytes(), &authResp1)
	token1 := authResp1["token"]
	assert.NotEmpty(t, token1)

	authReq2 := map[string]string{"username": "user2", "password": "pass2"}
	authJSON2, _ := json.Marshal(authReq2)
	req2, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON2))
	req2.Header.Set("Content-Type", "application/json")
	rec2 := httptest.NewRecorder()
	router.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	var authResp2 map[string]string
	json.Unmarshal(rec2.Body.Bytes(), &authResp2)
	token2 := authResp2["token"]
	assert.NotEmpty(t, token2)

	sendBody := map[string]interface{}{
		"toUser": "user2",
		"amount": 100,
	}
	sendJSON, _ := json.Marshal(sendBody)
	reqSend, _ := http.NewRequest("POST", "/api/sendCoin", bytes.NewBuffer(sendJSON))
	reqSend.Header.Set("Content-Type", "application/json")
	reqSend.Header.Set("Authorization", "Bearer "+token1)
	recSend := httptest.NewRecorder()
	router.ServeHTTP(recSend, reqSend)
	assert.Equal(t, http.StatusOK, recSend.Code)

	user1, err1 := shopUC.GetUserInfo("user1")
	user2, err2 := shopUC.GetUserInfo("user2")
	assert.NoError(t, err1)
	assert.NoError(t, err2)

	assert.Equal(t, 900, user1.Coins)
	assert.Equal(t, 1100, user2.Coins)
}

/*
TestIntegration_ConcurrentBuyItem tests concurrent purchase requests to the /api/buy/{item}
endpoint. It issues multiple simultaneous requests for buying an item and then verifies that
the user's coin balance and inventory reflect all successful purchases.
*/
func TestIntegration_ConcurrentBuyItem(t *testing.T) {
	router, shopUC, _ := setupRouter()

	authReq := map[string]string{"username": "buyer", "password": "buytest"}
	authJSON, _ := json.Marshal(authReq)
	reqAuth, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON))
	reqAuth.Header.Set("Content-Type", "application/json")
	recAuth := httptest.NewRecorder()
	router.ServeHTTP(recAuth, reqAuth)
	assert.Equal(t, http.StatusOK, recAuth.Code)
	var authResp map[string]string
	json.Unmarshal(recAuth.Body.Bytes(), &authResp)
	token := authResp["token"]
	assert.NotEmpty(t, token)

	const numPurchases = 20
	var wg sync.WaitGroup
	wg.Add(numPurchases)

	for i := 0; i < numPurchases; i++ {
		go func() {
			defer wg.Done()
			reqBuy, _ := http.NewRequest("GET", "/api/buy/socks", nil)
			reqBuy.Header.Set("Authorization", "Bearer "+token)
			recBuy := httptest.NewRecorder()
			router.ServeHTTP(recBuy, reqBuy)
			assert.Equal(t, http.StatusOK, recBuy.Code)
		}()
	}
	wg.Wait()

	user, err := shopUC.GetUserInfo("buyer")
	assert.NoError(t, err)
	expectedCoins := 1000 - numPurchases*10
	assert.Equal(t, expectedCoins, user.Coins)
	qty, exists := user.Inventory.Load("socks")
	assert.True(t, exists)
	assert.Equal(t, numPurchases, qty)
}

/*
TestIntegration_ConcurrentSendCoin tests concurrent coin transfer requests to the /api/sendCoin
endpoint. It performs multiple simultaneous transfers from one user to another and then verifies
that the final coin balances match the expected results.
*/
func TestIntegration_ConcurrentSendCoin(t *testing.T) {
	router, shopUC, _ := setupRouter()

	authReq1 := map[string]string{"username": "user1", "password": "pass1"}
	authJSON1, _ := json.Marshal(authReq1)
	req1, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON1))
	req1.Header.Set("Content-Type", "application/json")
	rec1 := httptest.NewRecorder()
	router.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)
	var authResp1 map[string]string
	json.Unmarshal(rec1.Body.Bytes(), &authResp1)
	token1 := authResp1["token"]
	assert.NotEmpty(t, token1)

	authReq2 := map[string]string{"username": "user2", "password": "pass2"}
	authJSON2, _ := json.Marshal(authReq2)
	req2, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(authJSON2))
	req2.Header.Set("Content-Type", "application/json")
	rec2 := httptest.NewRecorder()
	router.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	var authResp2 map[string]string
	json.Unmarshal(rec2.Body.Bytes(), &authResp2)
	token2 := authResp2["token"]
	assert.NotEmpty(t, token2)

	const numTransfers = 20
	const transferAmount = 50
	var wg sync.WaitGroup
	wg.Add(numTransfers)

	for i := 0; i < numTransfers; i++ {
		go func() {
			defer wg.Done()
			sendBody := map[string]interface{}{
				"toUser": "user2",
				"amount": transferAmount,
			}
			sendJSON, _ := json.Marshal(sendBody)
			reqSend, _ := http.NewRequest("POST", "/api/sendCoin", bytes.NewBuffer(sendJSON))
			reqSend.Header.Set("Content-Type", "application/json")
			reqSend.Header.Set("Authorization", "Bearer "+token1)
			recSend := httptest.NewRecorder()
			router.ServeHTTP(recSend, reqSend)
			assert.Equal(t, http.StatusOK, recSend.Code)
		}()
	}
	wg.Wait()

	user1, err1 := shopUC.GetUserInfo("user1")
	user2, err2 := shopUC.GetUserInfo("user2")
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	expectedUser1Balance := 1000 - numTransfers*transferAmount
	expectedUser2Balance := 1000 + numTransfers*transferAmount
	assert.Equal(t, expectedUser1Balance, user1.Coins)
	assert.Equal(t, expectedUser2Balance, user2.Coins)
}
