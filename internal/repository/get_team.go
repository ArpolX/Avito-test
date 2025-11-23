package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qUsers = `SELECT user_id, username, is_active FROM users WHERE team_name = $1`

func (r *RepositoryImpl) GetTeam(ctx context.Context, teamName string) (entity.Team, error) {
	t := entity.Team{
		TeamName: teamName,
	}

	rows, err := r.Postgres.DB.Query(ctx, qUsers, teamName)
	if err != nil {
		return entity.Team{}, fmt.Errorf("Ошибка в Query, метод GetTeam: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u entity.TeamMember
		err := rows.Scan(&u.UserId, &u.Username, &u.IsActive)
		if err != nil {
			return entity.Team{}, err
		}
		t.Members = append(t.Members, u)
	}

	return t, nil
}
