package controller

import (
	"github.com/Daerys/avito-shop/internal/usecase"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"go.uber.org/zap"
)

type serverImpl struct {
	shop.UnimplementedAvitoShopServer
	coinUsecase usecase.CoinUsecase
	itemUsecase usecase.ItemUsecase
	userUsecase usecase.UserUsecase

	sugar *zap.SugaredLogger
}

func NewServer(coinUsecase usecase.CoinUsecase, itemUsecase usecase.ItemUsecase, userUsecase usecase.UserUsecase, sugar *zap.SugaredLogger) *serverImpl {
	return &serverImpl{
		coinUsecase: coinUsecase,
		itemUsecase: itemUsecase,
		userUsecase: userUsecase,
		sugar:       sugar,
	}
}
