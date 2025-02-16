package gateway

import (
	"context"
	"net/http"

	"github.com/Daerys/avito-shop/config"
	"github.com/Daerys/avito-shop/pkg/generated/api/shop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// StartGateway запускает HTTP сервер, проксирующий запросы в gRPC.
func StartGateway(ctx context.Context, cfg *config.Config, sugar *zap.SugaredLogger) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := shop.RegisterAvitoShopHandlerFromEndpoint(ctx, mux, "localhost:"+cfg.GRPC.Port, opts); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":" + cfg.GRPC.GatewayPort,
		Handler: mux,
	}

	// Завершение работы HTTP сервера при отмене контекста
	go func() {
		<-ctx.Done()
		sugar.Info("Остановка gRPC-Gateway...")
		srv.Shutdown(ctx)
	}()

	sugar.Infof("gRPC-Gateway запущен на порту %s", cfg.GRPC.GatewayPort)
	return srv.ListenAndServe()
}
