package controller

import (
	"context"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository/mock"
	"github.com/Daerys/avito-shop/internal/usecase"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestAuth_ValidCredentials(t *testing.T) {
	users := make(map[string]*entity.User)
	uc := usecase.NewUserUsecase(mock.NewUserRepository(&users), "123123")
	server := &serverImpl{
		userUsecase: uc,
	}
	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}
	_, err := server.Auth(context.Background(), req)
	assert.NoError(t, err)

	_, err = server.Auth(context.Background(), req)
	assert.NoError(t, err)
}

func TestAuth_InvalidPassword(t *testing.T) {
	users := make(map[string]*entity.User)
	uc := usecase.NewUserUsecase(mock.NewUserRepository(&users), "123123")
	logger, err := zap.NewProduction()
	sugar := logger.Sugar()
	server := &serverImpl{
		userUsecase: uc,
		sugar:       sugar,
	}
	req := &shop.AuthRequest{
		Username: "user",
		Password: "123",
	}
	_, err = server.Auth(context.Background(), req)
	assert.NoError(t, err)

	req = &shop.AuthRequest{
		Username: "user",
		Password: "122",
	}
	_, err = server.Auth(context.Background(), req)
	assert.Error(t, err)
}

func TestAuth_InvalidCredentials(t *testing.T) {
	users := make(map[string]*entity.User)
	uc := usecase.NewUserUsecase(mock.NewUserRepository(&users), "123123")
	logger, err := zap.NewProduction()
	sugar := logger.Sugar()
	server := &serverImpl{
		userUsecase: uc,
		sugar:       sugar,
	}
	req := &shop.AuthRequest{
		Username: "user1",
		Password: "",
	}
	_, err = server.Auth(context.Background(), req)
	assert.Error(t, err)

	req = &shop.AuthRequest{
		Username: "",
		Password: "122",
	}
	_, err = server.Auth(context.Background(), req)
	assert.Error(t, err)
}
