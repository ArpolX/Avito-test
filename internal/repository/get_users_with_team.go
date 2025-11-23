package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qGetUsersWithTeam = `
		SELECT user_id, username, team_name, is_active
		FROM users WHERE team_name = $1
	`

func (r *RepositoryImpl) GetUsersWithTeam(ctx context.Context, teamName string) ([]entity.User, error) {
	rows, err := r.Postgres.DB.Query(ctx, qGetUsersWithTeam, teamName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в Query, метод GetUsersWithTeam: %w", err)
	}
	defer rows.Close()

	var res []entity.User

	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.UserId, &u.Username, &u.TeamName, &u.IsActive)
		if err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}
