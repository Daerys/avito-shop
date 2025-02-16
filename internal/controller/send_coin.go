package controller

import (
	"context"
	"strconv"

	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverImpl) SendCoin(ctx context.Context, req *shop.SendCoinRequest) (*shop.SendCoinResponse, error) {
	err := s.checkUser(ctx, req.FromUser)
	if err != nil {
		return nil, err
	}

	// Proceed with sending coins.
	transaction, err := s.coinUsecase.SendCoin(ctx, req.ToUser, req.FromUser, req.Amount)
	if err != nil {
		// 500 Internal Server Error: An error occurred while processing the coin transfer.
		return nil, status.Errorf(codes.Internal, "failed to send coins: %v", err)
	}

	// Retrieve the updated user information.
	user, err := s.userUsecase.GetUser(ctx, req.FromUser)
	if err != nil {
		// 500 Internal Server Error: Unable to fetch user details.
		return nil, status.Errorf(codes.Internal, "failed to retrieve user info: %v", err)
	}

	// Return successful response.
	return &shop.SendCoinResponse{
		RemainingCoins: int32(user.Coins),
		TransactionId:  strconv.Itoa(int(transaction.ID)),
	}, nil
}
