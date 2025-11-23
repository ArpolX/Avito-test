package postgres

import (
	"avito-test/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func NewConnectDB(ctx context.Context, cfg config.Config, log *zap.Logger) (*Postgres, error) {
	log.Info("Connecting to the DB...")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.UserDB, cfg.DB.PasswordDB, cfg.DB.HostDB, cfg.DB.PortDB, cfg.DB.NameDB)

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(ctx, time.Second*8)
	defer cancel()

	count := 0

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("DB connection cancellation %w", ctx.Err())
		case <-ticker.C:
			pool, err := pgxpool.New(ctx, dsn)
			if err != nil {
				log.Error("connection pool creation error attempt", zap.Error(err))
				count++
				break
			}
			err = pool.Ping(ctx)
			if err != nil {
				log.Error("error PING attempt", zap.Error(err))
				pool.Close()
				count++
				break
			}
			log.Info("Successful connection to DB")
			return &Postgres{
				DB: pool,
			}, nil
		}
	}

}
