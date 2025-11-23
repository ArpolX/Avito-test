package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qGetUser = `
		SELECT user_id, username, team_name, is_active
		FROM users WHERE user_id = $1
	`

func (r *RepositoryImpl) GetUser(ctx context.Context, userId string) (entity.User, error) {
	var u entity.User

	err := r.Postgres.DB.QueryRow(ctx, qGetUser, userId).Scan(
		&u.UserId, &u.Username, &u.TeamName, &u.IsActive,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("Ошибка в exec, метод GetUser: %w", err)
	}
	return u, nil
}
