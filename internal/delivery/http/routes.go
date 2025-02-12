package http

import (
	"avito-shop/internal/usecase/shop"
	"github.com/gin-gonic/gin"
)

/*
RegisterRoutes registers all HTTP endpoints for the Avito Shop API.
It sets up the authentication endpoint without JWT and a group of endpoints that require a valid JWT.
*/
func RegisterRoutes(router *gin.Engine, usecase *usecase.ShopUsecase, jwtSecret string) {
	router.POST("/api/auth", AuthHandler(usecase, jwtSecret))

	authGroup := router.Group("/api")
	authGroup.Use(JWTMiddleware(jwtSecret))
	{
		authGroup.GET("/info", NewHandler(usecase).Info)
		authGroup.POST("/sendCoin", NewHandler(usecase).SendCoin)
		authGroup.GET("/buy/:item", NewHandler(usecase).BuyItem)
	}
}
