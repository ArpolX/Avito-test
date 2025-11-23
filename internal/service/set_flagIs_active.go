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

func (s *ServiceImpl) SetFlagIsActive(ctx context.Context, setActive entity.SetIsActiveRequest) (entity.User, error) {
	if err := s.Repo.UpdateUser(ctx, setActive.UserID, setActive.IsActive); err != nil {
		s.Log.Error("Ошибка при обновлении юзера, метод SetFlagIsActive", zap.Error(err))
		return entity.User{}, fmt.Errorf("Ошибка при обновлении юзера, метод SetFlagIsActive: %w", err)
	}

	user, err := s.Repo.GetUser(ctx, setActive.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка при получении юзера, метод SetFlagIsActive", zap.Error(err))
		return entity.User{}, fmt.Errorf("Ошибка при получении юзера, метод SetFlagIsActive: %w", err)
	}

	return user, nil
}
