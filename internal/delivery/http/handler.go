package http

import (
	"log"
	"net/http"

	"avito-shop/internal/entity"
	"avito-shop/internal/usecase/shop"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.ShopUsecase
}

// NewHandler creates and returns a new Handler instance with the given ShopUsecase.
func NewHandler(u *usecase.ShopUsecase) *Handler {
	return &Handler{usecase: u}
}

// Info handles GET /api/info requests.
func (h *Handler) Info(c *gin.Context) {
	username := c.GetString("username")
	log.Printf("Info: request for user info, username: %s", username)
	if username == "" {
		log.Printf("Info: unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "unauthorized"})
		return
	}
	user, err := h.usecase.GetUserInfo(username)
	if err != nil {
		log.Printf("Info: error retrieving user info for %s: %v", username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed to get user info"})
		return
	}
	transactions, err := h.usecase.GetTransactions(username)
	if err != nil {
		log.Printf("Info: error retrieving transactions for %s: %v", username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed to get transactions"})
		return
	}

	var received []gin.H
	var sent []gin.H
	for _, tx := range transactions {
		if tx.ToUser == username && tx.Type == entity.TransactionTypeSend {
			received = append(received, gin.H{"fromUser": tx.FromUser, "amount": tx.Amount})
		}
		if tx.FromUser == username && tx.Type == entity.TransactionTypeSend {
			sent = append(sent, gin.H{"toUser": tx.ToUser, "amount": tx.Amount})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"coins":     user.Coins,
		"inventory": user.Inventory,
		"coinHistory": gin.H{
			"received": received,
			"sent":     sent,
		},
	})
	log.Printf("Info: response sent for user %s", username)
}

// SendCoin handles POST /api/sendCoin requests.
func (h *Handler) SendCoin(c *gin.Context) {
	username := c.GetString("username")
	log.Printf("SendCoin: request received from user: %s", username)
	if username == "" {
		log.Printf("SendCoin: unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "unauthorized"})
		return
	}
	var req struct {
		ToUser string `json:"toUser"`
		Amount int    `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("SendCoin: invalid request payload from user %s: %v", username, err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": "invalid request"})
		return
	}
	if err := h.usecase.SendCoins(username, req.ToUser, req.Amount); err != nil {
		log.Printf("SendCoin: error sending coins from %s to %s: %v", username, req.ToUser, err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coins sent successfully"})
	log.Printf("SendCoin: coins sent successfully from %s to %s", username, req.ToUser)
}

// BuyItem handles GET /api/buy/:item requests.
func (h *Handler) BuyItem(c *gin.Context) {
	username := c.GetString("username")
	log.Printf("BuyItem: request received from user: %s", username)
	if username == "" {
		log.Printf("BuyItem: unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "unauthorized"})
		return
	}
	itemName := c.Param("item")
	if itemName == "" {
		log.Printf("BuyItem: item not specified by user %s", username)
		c.JSON(http.StatusBadRequest, gin.H{"errors": "item not specified"})
		return
	}
	if err := h.usecase.BuyItem(username, itemName); err != nil {
		log.Printf("BuyItem: error buying item %s for user %s: %v", itemName, username, err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "item purchased successfully"})
	log.Printf("BuyItem: item %s purchased successfully by user %s", itemName, username)
}
