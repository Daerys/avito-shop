package controller

import (
	"context"
	"strings"

	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverImpl) Auth(ctx context.Context, req *shop.AuthRequest) (*shop.AuthResponse, error) {
	// Validate input
	if req.Username == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username and password are required")
	}

	// Attempt to register or authenticate the user.
	token, err := s.userUsecase.RegisterOrAuthenticate(ctx, req.Username, req.Password)
	if err != nil {
		s.sugar.Errorf("[ERROR] Couldn't authenticate user: %v", err)
		// Check error details to return a more specific status code.
		if strings.Contains(err.Error(), "invalid password") {
			// 401 Unauthorized: Invalid credentials.
			return nil, status.Error(codes.Unauthenticated, "invalid username or password")
		} else if strings.Contains(err.Error(), "db error") {
			// 500 Internal Server Error: Database error.
			return nil, status.Errorf(codes.Internal, "internal server error")
		}
		// Default to 400 Bad Request for other errors.
		return nil, status.Errorf(codes.InvalidArgument, "bad request: %v", err)
	}

	// Return a successful authentication response.
	return &shop.AuthResponse{Token: token}, nil
}
