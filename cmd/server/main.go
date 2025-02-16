package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Daerys/avito-shop/config"
	"github.com/Daerys/avito-shop/internal/app"
	"go.uber.org/zap"
)

func main() {
	// Создаем zap-логгер (Production-конфигурация)
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Unable to create logger: " + err.Error())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// Загружаем конфигурацию
	cfg, err := config.NewConfig()
	if err != nil {
		sugar.Fatalf("Error while loading config: %v", err)
	}

	// Создаем базовый контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов завершения
	setupSignalHandler(cancel, sugar)

	// Запускаем приложение (gRPC сервер и Gateway)
	if err := app.Start(ctx, cfg, sugar); err != nil {
		sugar.Fatalf("Application start error: %v", err)
	}

	// Ожидаем отмены контекста
	<-ctx.Done()
	sugar.Info("graceful shutdown")
}

// setupSignalHandler отменяет контекст при получении сигнала завершения.
func setupSignalHandler(cancel func(), sugar *zap.SugaredLogger) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		sugar.Infof("syscall: %v", sig)
		cancel()
	}()
}
