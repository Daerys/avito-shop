package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Daerys/avito-shop/config"
	"github.com/Daerys/avito-shop/internal/controller"
	"github.com/Daerys/avito-shop/internal/repository/postgres"
	"github.com/Daerys/avito-shop/internal/usecase"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func StartGRPCServer(ctx context.Context, cfg *config.Config, sugar *zap.SugaredLogger) error {
	dsn := cfg.PG.URL
	if cfg.PG.URL == "" {
		dsn = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			cfg.PG.Host, cfg.PG.Port, cfg.PG.DB, cfg.PG.User, cfg.PG.Password)

	}

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		sugar.Errorf("Unable to connect to data base by url: %s", dsn)
	}

	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWT.Secret)
	coinRepo := postgres.NewCoinRepository(db)
	coinUsecase := usecase.NewCoinUsecase(coinRepo)
	itemRepo := postgres.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepo)
	grpcServer := grpc.NewServer()
	shop.RegisterAvitoShopServer(grpcServer, controller.NewServer(coinUsecase, itemUsecase, userUsecase, sugar))

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		return err
	}

	// Завершение работы сервера при отмене контекста
	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
		sugar.Info("gRPC server stopped")
	}()

	sugar.Infof("gRPC server started on port: %s", cfg.GRPC.Port)
	return grpcServer.Serve(lis)
}
