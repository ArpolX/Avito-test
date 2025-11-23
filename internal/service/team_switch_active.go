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

func (s *ServiceImpl) TeamSwitchActive(ctx context.Context, teamName string) (entity.Team, error) {
	users, err := s.Repo.GetUsersWithTeam(ctx, teamName)
	if err != nil {
		s.Log.Error("Ошибка при получении юзеров команды, метод TeamSwitchActive", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при получении юзеров команды, метод TeamSwitchActive: %w", err)
	}

	for _, user := range users {
		if user.IsActive == true {
			if err := s.Repo.UpdateUser(ctx, user.UserId, false); err != nil {
				s.Log.Error("Ошибка при изменении активности юзеров, метод TeamSwitchActive", zap.Error(err))
				return entity.Team{}, fmt.Errorf("Ошибка при изменении активности юзеров, метод TeamSwitchActive: %w", err)
			}
		}
	}

	// обработка здесь, потому что когда делаем query и собираем массив, ошибка ErrNoRows не возвращается
	team, err := s.Repo.GetTeam(ctx, teamName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Team{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка при получении команды, метод TeamSwitchActive", zap.Error(err))
		return entity.Team{}, fmt.Errorf("Ошибка при получении команды, метод TeamSwitchActive: %w", err)
	}

	return team, nil
}
