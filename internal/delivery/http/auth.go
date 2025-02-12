package http

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"avito-shop/internal/entity"
	"avito-shop/internal/usecase/shop"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var expTime = time.Now().Add(24 * time.Hour).Unix()

// AuthHandler handles user authentication and JWT token generation.
func AuthHandler(u *usecase.ShopUsecase, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("AuthHandler: invalid request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": "invalid request"})
			return
		}
		log.Printf("AuthHandler: request received for username: %s", req.Username)

		user, err := u.GetUserInfo(req.Username)
		if err != nil {
			// Если пользователь не найден – пытаемся создать нового
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("AuthHandler: user not found, creating new user: %s", req.Username)
				if _, ok := createNewUser(c, req, user, u); ok {
					return
				}
			} else {
				log.Printf("AuthHandler: error retrieving user info for %s: %v", req.Username, err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		} else {
			if checkPassword(c, user, req) {
				log.Printf("AuthHandler: invalid password for user: %s", req.Username)
				return
			}
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      expTime,
		})
		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			log.Printf("AuthHandler: failed to generate token for user %s: %v", req.Username, err)
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		log.Printf("AuthHandler: successful auth for user %s, duration: %s", req.Username, time.Since(start))
	}
}

// createNewUser creates a new user with the provided credentials.
func createNewUser(c *gin.Context, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}, user *entity.User, u *usecase.ShopUsecase) (*entity.User, bool) {
	log.Printf("createNewUser: creating new user: %s", req.Username)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("createNewUser: failed to hash password for %s: %v", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed to hash password"})
		return nil, true
	}

	newUser := &entity.User{
		Username:     req.Username,
		Coins:        1000,
		Inventory:    &sync.Map{},
		PasswordHash: string(hashedPass),
	}
	user, err = u.CreateUser(newUser)
	if err != nil {
		log.Printf("createNewUser: failed to create user %s: %v", req.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed to create user"})
		return nil, true
	}
	log.Printf("createNewUser: user %s created successfully", req.Username)
	return user, false
}

// checkPassword compares the provided password with the stored password hash.
func checkPassword(c *gin.Context, user *entity.User, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Printf("checkPassword: invalid credentials for user %s", req.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "invalid credentials"})
		return true
	}
	return false
}
