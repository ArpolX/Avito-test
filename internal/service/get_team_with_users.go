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

func (s *ServiceImpl) GetTeamWithUsers(ctx context.Context, teamName string) (entity.Team, error) {
	tOrig, err := s.Repo.GetTeam(ctx, teamName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Team{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка при получении команды, метод CreateTeamWithUsers", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при получении команды, метод CreateTeamWithUsers: %w", err)
	}

	return tOrig, nil
}
