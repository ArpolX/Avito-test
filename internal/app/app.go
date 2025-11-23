package app

import (
	"avito-test/internal/config"
	"avito-test/internal/controller"
	"avito-test/internal/infrastructure/database/postgres"
	"avito-test/internal/infrastructure/logger"
	"avito-test/internal/repository"
	"avito-test/internal/route"
	"avito-test/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const httpMaxRetries = 3

func RunApp(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to read .env, метод RunApp: %w", err)
	}

	log, err := logger.InitLogger()
	if err != nil {
		return fmt.Errorf("failed to start zap.logger, метод RunApp: %w", err)
	}

	postgresDB, err := postgres.NewConnectDB(ctx, cfg, log)
	if err != nil {
		return fmt.Errorf("failed to start database(postgreSQL), метод RunApp: %w", err)
	}

	if err := RunRestServer(ctx, cfg, log, postgresDB); err != nil {
		log.Error("failed to start rest server, метод RunApp", zap.Error(err))
		return fmt.Errorf("failed to start rest server, метод RunApp: %w", err)
	}

	return nil
}

func RunRestServer(ctx context.Context, cfg config.Config, log *zap.Logger, postgres *postgres.Postgres) error {
	log.Info("starting rest server")

	addr := cfg.ServerHost + ":" + cfg.ServerPort

	repo := repository.NewRepositoryImpl(postgres, log)
	srv := service.NewServiceImpl(log, repo)
	ctrl := controller.NewControllerImpl(cfg, log, srv)

	r := route.Handlers(&ctrl)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	for i := 0; i < httpMaxRetries; i++ {
		log.Info("Attempting to start HTTP server", zap.String("method", "RunRestServer"), zap.Int("attempt", i+1), zap.String("addr", addr))

		if err := server.ListenAndServe(); err != nil {
			log.Warn("Failed to start HTTP server", zap.String("method", "RunRestServerWithGraceful"), zap.Int("attempt", i+1), zap.Error(err))
		} else {
			log.Info("server is running")
			break
		}
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Info("Shutting down gracefully")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to during server shutdown, метод RunRestServer: %w", err)
	}

	return nil
}
