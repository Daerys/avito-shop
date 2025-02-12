package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"

	"avito-shop/internal/delivery/http"
	"avito-shop/internal/repository/postgres"
	"avito-shop/internal/usecase/shop"
)

func main() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	serverPort := os.Getenv("SERVER_PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "d1f54313ed1d00ed4140bf5a9b93ca45960559aab7f6276a8632ab099e29c33af6da4a1f9a6e3b774f9597fcb2047a69acf96552621db562998c81d4a3c078f27a08537980d2e7288f0351fb89a74ac94e4b7a22851f65c7c57229bd9f7eb90514f91570500f31f318c0528d85a819718a8385f1054a58d76c9aaf0472dec8cb0bec4125b10af73439177696fc25459d6a7c3e58150045c5873560ca84ad037af0f37cf4757b2bccfb3753e85eaf219c11bbd85f0c43b19083415bd5eec7d00f2ceda0dfc2fffc0b5497f88ab1006f1c48639b69bbea8c1c0e8568076aed7ca3d505617af01dd2b3bedab1addb5cd6c3cc94f5230a66923aec99b3e753950ce5" // значение по умолчанию
	}
	if serverPort == "" {
		serverPort = "8080"
	}

	psqlInfo := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	userRepo := postgres.NewUserRepository(db)
	transactionRepo := postgres.NewTransactionRepository(db)
	itemRepo := postgres.NewItemRepository(db)

	shopUsecase := usecase.NewShopUsecase(userRepo, transactionRepo, itemRepo)

	router := gin.Default()
	http.RegisterRoutes(router, shopUsecase, jwtSecret)

	log.Printf("Server running at http://localhost:%s", serverPort)
	router.Run(":" + serverPort)
}
