package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *serverImpl) BuyItem(ctx context.Context, req *shop.BuyItemRequest) (*shop.BuyItemResponse, error) {
	// Validate JWT and get the username
	userName, err := s.checkJWT(ctx)
	if err != nil {
		return nil, err
	}

	// Retrieve the item by its name.
	item, err := s.itemUsecase.GetItem(ctx, req.GetItem())
	if err != nil {
		// If the item is not found, consider it an invalid request (400).
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.InvalidArgument, "Couldn't find item: "+req.GetItem())
		}
		// Otherwise, it's a database/internal error.
		return nil, status.Error(codes.Internal, "Database error while retrieving item: "+err.Error())
	}

	// Retrieve the user by username.
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

	// Remove coins from the user.
	err = s.coinUsecase.RemoveCoins(ctx, user, item.Price)
	if err != nil {
		// If the error indicates insufficient balance, return 400 (Invalid Request).
		if strings.Contains(err.Error(), "insufficient balance") {
			return nil, status.Error(codes.InvalidArgument, "Insufficient coins")
		}
		// Otherwise, it's an internal error.
		return nil, status.Error(codes.Internal, "Database error while removing coins: "+err.Error())
	}

	// Add the item to the user's inventory.
	err = s.itemUsecase.AddItem(ctx, user, item)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error while updating inventory: "+err.Error())
	}

	items, err := s.itemUsecase.GetInventory(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error while retrieving inventory: "+err.Error())
	}

	return &shop.BuyItemResponse{
		RemainingCoins: int32(user.Coins),
		Items:          entityToGRPCInventoryItemSlice(items),
	}, nil
}
