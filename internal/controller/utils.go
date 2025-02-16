package controller

import (
	"context"
	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (s *serverImpl) checkJWT(ctx context.Context) (string, error) {
	// Validate the JWT token from the request metadata.
	token, err := s.userUsecase.ParseToken(ctx)
	if err != nil {
		// 401 Unauthorized: Token is missing or invalid.
		return "", status.Error(codes.Unauthenticated, "unauthenticated: "+err.Error())
	}

	// Retrieve the username from the token claims.
	username, err := s.userUsecase.GetUserFromToken(ctx, token)
	if err != nil {
		// 401 Unauthorized: Unable to extract valid user information from the token.
		return "", status.Error(codes.Unauthenticated, "unauthenticated: "+err.Error())
	}
	return username, nil
}

func (s *serverImpl) checkUser(ctx context.Context, userReq string) error {
	user, err := s.checkJWT(ctx)
	if err != nil {
		return err
	}
	// Ensure that the token belongs to the user performing the operation.
	if userReq != user {
		// 403 Permission Denied: Users are not allowed to perform operations on behalf of another user.
		return status.Error(codes.PermissionDenied, "you are not allowed to perform operations on behalf of another user")
	}
	return nil
}

func entityToGRPCInventoryItemSlice(items []entity.InventoryItem) []*shop.InventoryItem {
	resultItems := make([]*shop.InventoryItem, 0, len(items))
	for _, invItem := range items {
		resultItems = append(resultItems, &shop.InventoryItem{
			Type:     invItem.ItemType,
			Quantity: int32(invItem.Quantity),
		})
	}
	return resultItems
}

func entityToGRPCCoinHTransactionSlice(items []entity.CoinTransaction, user *entity.User) *shop.CoinHistory {
	receivedTransactions := make([]*shop.Received, 0)
	sentTransactions := make([]*shop.Sent, 0)

	for _, transaction := range items {
		if transaction.FromUserID == user.ID {
			sentTransactions = append(sentTransactions, &shop.Sent{
				ToUser: strconv.FormatInt(transaction.ToUserID, 10),
				Amount: int32(transaction.Amount),
			})
		} else if transaction.ToUserID == user.ID {
			receivedTransactions = append(receivedTransactions, &shop.Received{
				FromUser: strconv.FormatInt(transaction.FromUserID, 10),
				Amount:   int32(transaction.Amount),
			})
		}
	}
	return &shop.CoinHistory{
		Received: receivedTransactions,
		Sent:     sentTransactions,
	}
}
