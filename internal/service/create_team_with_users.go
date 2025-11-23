package service

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (s *ServiceImpl) CreateTeamWithUsers(ctx context.Context, team entity.Team) (entity.Team, error) {
	_, err := s.Repo.GetTeam(ctx, team.TeamName)
	if err == nil {
		return entity.Team{}, fmt.Errorf("%w", Error.TEAM_EXISTS)
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		s.Log.Error("Ошибка при получении команды, метод CreateTeamWithUsers", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при получении команды, метод CreateTeamWithUsers: %w", err)
	}

	if err := s.Repo.CreateTeamWithUsers(ctx, team.TeamName, team.Members); err != nil {
		s.Log.Error("Ошибка при создании команды с юзерами, метод CreateTeamWithUsers:", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при создании команды с юзерами, метод CreateTeamWithUsers:: %w", err)
	}

	tOrig, err := s.Repo.GetTeam(ctx, team.TeamName)
	if err != nil {
		s.Log.Error("Ошибка при получении команды, метод CreateTeamWithUsers", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при получении команды, метод CreateTeamWithUsers: %w", err)
	}

	return tOrig, nil
}
