package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Daerys/avito-shop/internal/entity"
	"github.com/Daerys/avito-shop/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

// Error constants.
var (
	ErrNoMetadata              = errors.New("no metadata in request")
	ErrNoAuthToken             = errors.New("authorization token not provided")
	ErrInvalidTokenFormat      = errors.New("invalid token format")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrFailedToGenerateToken   = errors.New("failed to generate token")
	ErrCouldNotParseClaims     = errors.New("could not parse claims")
	ErrUsernameNotFound        = errors.New("username not found in token")
)

// UserUsecase defines user-related operations.
type UserUsecase interface {
	GetUser(ctx context.Context, userName string) (*entity.User, error)
	RegisterOrAuthenticate(ctx context.Context, username, password string) (string, error)
	GetUserFromToken(ctx context.Context, token *jwt.Token) (string, error)
	ParseToken(ctx context.Context) (*jwt.Token, error)
}

type userUsecaseImpl struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewUserUsecase(repo repository.UserRepository, jwtSecret string) UserUsecase {
	return &userUsecaseImpl{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (u *userUsecaseImpl) GetUserFromToken(ctx context.Context, token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrCouldNotParseClaims
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", ErrUsernameNotFound
	}

	return username, nil
}

func (u *userUsecaseImpl) ParseToken(ctx context.Context) (*jwt.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNoMetadata
	}

	// Retrieve the Authorization header.
	authHeaders, ok := md["authorization"]
	if !ok || len(authHeaders) == 0 {
		return nil, ErrNoAuthToken
	}

	// Expect format "Bearer <token>"
	authHeader := authHeaders[0]
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, ErrInvalidTokenFormat
	}
	tokenString := parts[1]

	// Parse and validate the JWT token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(u.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}
	return token, nil
}

func (u *userUsecaseImpl) GetUser(ctx context.Context, userName string) (*entity.User, error) {
	return u.repo.GetUserByUsername(ctx, userName)
}

func (u *userUsecaseImpl) RegisterOrAuthenticate(ctx context.Context, username, password string) (string, error) {
	hashedPass, err := hashPassword(password)
	if err != nil {
		return "", err
	}

	user, err := u.repo.RegisterOrAuthenticate(ctx, username, hashedPass)
	if err != nil {
		return "", fmt.Errorf("db error: %w", err)
	}

	if err := checkPassword(user.PasswordHash, password); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return u.createJWT(user)
}

func (u *userUsecaseImpl) createJWT(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", ErrFailedToGenerateToken
	}
	return tokenString, nil
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}
