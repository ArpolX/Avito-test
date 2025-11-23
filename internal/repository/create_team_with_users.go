package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const qGetTeam = `INSERT INTO teams (team_name) VALUES ($1)`

const qGetUsers = `INSERT INTO users (user_id, username, is_active, team_name) 
      			VALUES ($1, $2, $3, $4)`

func (r *RepositoryImpl) CreateTeamWithUsers(ctx context.Context, teamName string, users []entity.TeamMember) error {
	tx, err := r.Postgres.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("Ошибка в начале транзакции, метод CreateTeamWithUsers: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, qGetTeam, teamName)
	if err != nil {
		return fmt.Errorf("Ошибка в exec, метод CreateTeamWithUsers: %w", err)
	}

	for _, u := range users {
		_, err := tx.Exec(ctx,
			qGetUsers,
			u.UserId, u.Username, u.IsActive, teamName,
		)
		if err != nil {
			return fmt.Errorf("Ошибка в exec, метод CreateTeamWithUsers: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("Ошибка в commit транзакции, метод CreateTeamWithUsers: %w", err)
	}

	return nil
}
