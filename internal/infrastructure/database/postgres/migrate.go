package postgres

import (
	"avito-test/internal/config"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
	"go.uber.org/zap"
)

func RunMigrations(ctx context.Context, cfg config.Config, log *zap.Logger) error {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.HostDB, cfg.DB.PortDB, cfg.DB.UserDB, cfg.DB.PasswordDB, cfg.DB.NameDB)

	db, err := sql.Open("pgx", connect)
	if err != nil {
		log.Error("opening error pgx", zap.Error(err))
		return fmt.Errorf("opening error pgx: %w", err)
	}
	defer db.Close()
	if err := goose.Up(db, "./app/migrations"); err != nil {
		return fmt.Errorf("error starting migration: %w", err)
	}

	log.Info("Successful Migration")
	return nil
}
