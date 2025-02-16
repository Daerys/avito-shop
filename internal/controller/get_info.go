package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverImpl) GetInfo(ctx context.Context, req *shop.InfoRequest) (*shop.InfoResponse, error) {
	// Validate JWT and get the username
	userName, err := s.checkJWT(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.userUsecase.GetUser(ctx, userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || user == nil {
			return nil, status.Error(codes.InvalidArgument, "Couldn't find user: "+userName)
		}
		return nil, status.Error(codes.Internal, "Database error while retrieving user: "+err.Error())
	}
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "User not found: "+userName)
	}

	its, err := s.itemUsecase.GetInventory(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error while retrieving inventory: "+err.Error())
	}
	items := entityToGRPCInventoryItemSlice(its)

	coinHistory, err := s.coinUsecase.GetHistory(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error while retrieving coin history: "+err.Error())
	}
	return &shop.InfoResponse{
		Coins:       int32(user.Coins),
		Inventory:   items,
		CoinHistory: entityToGRPCCoinHTransactionSlice(coinHistory, user),
	}, nil
}
