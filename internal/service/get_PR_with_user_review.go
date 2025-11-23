package service

import (
	"avito-test/internal/entity"
	"context"
	"fmt"

	"go.uber.org/zap"
)

func (s *ServiceImpl) GetPRWithUserReview(ctx context.Context, userId string) ([]entity.PullRequestShort, error) {
	prs, err := s.Repo.GetPRSWithUser(ctx, userId)
	if err != nil {
		s.Log.Error("Ошибка при получении prs, метод GetPRWithUserReview", zap.Error(err))
		return nil, fmt.Errorf("Ошибка при получении prs, метод GetPRWithUserReview: %w", err)
	}

	return prs, nil
}
