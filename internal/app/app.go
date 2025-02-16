package app

import (
	"context"
	"sync"

	"github.com/Daerys/avito-shop/config"
	"github.com/Daerys/avito-shop/internal/app/gateway"
	"github.com/Daerys/avito-shop/internal/app/server"
	"go.uber.org/zap"
)

// Start запускает gRPC сервер и gRPC-Gateway параллельно.
func Start(ctx context.Context, cfg *config.Config, sugar *zap.SugaredLogger) error {
	var wg sync.WaitGroup

	// Запуск gRPC сервера
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.StartGRPCServer(ctx, cfg, sugar); err != nil {
			sugar.Errorf("gRPC server ERROR: %v", err)
		}
	}()

	// Запуск gRPC-Gateway
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := gateway.StartGateway(ctx, cfg, sugar); err != nil {
			sugar.Errorf("gRPC-Gateway ERROR: %v", err)
		}
	}()

	wg.Wait()
	return nil
}
