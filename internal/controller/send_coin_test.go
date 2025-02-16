package controller

import (
	"context"
	"testing"

	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository/mock"
	"github.com/Daerys/avito-shop/internal/usecase"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestSendCoin_Success(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()
	req := &shop.AuthRequest{
		Username: "alice",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)
	req = &shop.AuthRequest{
		Username: "bob",
		Password: "123",
	}

	_, _ = server.Auth(ctx, req)

	resp, err := server.SendCoin(ctx, &shop.SendCoinRequest{
		FromUser: "alice",
		ToUser:   "bob",
		Amount:   50,
	})

	assert.NoError(t, err)
	assert.Equal(t, int32(950), resp.RemainingCoins)
}

func TestSendCoin_InsufficientFunds(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	server := &serverImpl{
		userUsecase: userUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()
	req := &shop.AuthRequest{
		Username: "alice",
		Password: "123",
	}

	tokenString, _ := server.Auth(ctx, req)
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + tokenString.Token,
	})
	ctx = metadata.NewIncomingContext(ctx, md)
	req = &shop.AuthRequest{
		Username: "bob",
		Password: "123",
	}

	_, _ = server.Auth(ctx, req)

	_, err := server.SendCoin(ctx, &shop.SendCoinRequest{
		FromUser: "alice",
		ToUser:   "bob",
		Amount:   50000,
	})

	assert.Error(t, err)
}

func TestSendCoin_Unauthorized(t *testing.T) {
	users := make(map[string]*entity.User)
	userRepo := mock.NewUserRepository(&users)
	coinRepo := mock.NewCoinRepository(&users)
	userUsecase := usecase.NewUserUsecase(userRepo, "123123")
	coinUsecase := usecase.NewCoinUsecase(coinRepo)

	server := &serverImpl{
		userUsecase: userUsecase,
		coinUsecase: coinUsecase,
	}

	ctx := context.Background()

	_, err := server.SendCoin(ctx, &shop.SendCoinRequest{
		FromUser: "alice",
		ToUser:   "bob",
		Amount:   50,
	})

	assert.Error(t, err)
}
