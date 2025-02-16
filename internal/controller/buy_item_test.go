package controller

import (
	"context"
	"github.com/Daerys/avito-shop/internal/entity"
	"google.golang.org/grpc/metadata"
	"testing"

	"github.com/Daerys/avito-shop/internal/repository/mock"
	"github.com/Daerys/avito-shop/internal/usecase"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"github.com/stretchr/testify/assert"
)

func TestBuyItem_Success(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	itemRepo := mock.NewItemRepository()
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		itemUsecase: itemUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()

	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	resp, err := server.BuyItem(ctx, &shop.BuyItemRequest{Item: "t-shirt"})
	assert.NoError(t, err)
	assert.Equal(t, int32(840), resp.RemainingCoins)
}

func TestBuyItem_ItemNotFound(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	itemRepo := mock.NewItemRepository()
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		itemUsecase: itemUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()

	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)
	_, err := server.BuyItem(ctx, &shop.BuyItemRequest{Item: "non-existent"})
	assert.Error(t, err)
}

func TestBuyItem_InsufficientFunds(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	itemRepo := mock.NewItemRepository()
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		itemUsecase: itemUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()

	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := server.BuyItem(ctx, &shop.BuyItemRequest{Item: "pink-hoody"})
	assert.NoError(t, err)

	_, err = server.BuyItem(ctx, &shop.BuyItemRequest{Item: "pink-hoody"})
	assert.Error(t, err)
}

func TestBuyItem_CheckInventory(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	itemRepo := mock.NewItemRepository()
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		itemUsecase: itemUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()

	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)

	resp, err := server.BuyItem(ctx, &shop.BuyItemRequest{Item: "t-shirt"})
	assert.NoError(t, err)
	inventory := resp.Items
	assert.Len(t, inventory, 1)
	assert.Equal(t, "t-shirt", inventory[0].Type)
	assert.Equal(t, 1, int(inventory[0].Quantity))
}
