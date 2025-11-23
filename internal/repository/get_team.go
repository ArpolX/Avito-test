package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qTeam = `SELECT team_name FROM teams WHERE team_name = $1`

const qUsers = `SELECT user_id, username, is_active FROM users WHERE team_name = $1`

func (r *RepositoryImpl) GetTeam(ctx context.Context, teamName string) (entity.Team, error) {
	var t entity.Team

	err := r.Postgres.DB.QueryRow(ctx, qTeam, teamName).Scan(&t.TeamName)
	if err != nil {
		return entity.Team{}, fmt.Errorf("Ошибка в Query, метод GetTeam: %w", err)
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
