package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

func (s *ServiceImpl) AmountPROpen(ctx context.Context) (int, error) {
	prs, err := s.Repo.GetAllPR(ctx)
	if err != nil {
		s.Log.Error("Ошибка при получении prs, метод AmountPROpen", zap.Error(err))
		return 0, fmt.Errorf("Ошибка при получении prs, метод AmountPROpen: %w", err)
	}

	count := 0
	for _, pr := range prs {
		if pr.Status == "Open" {
			count++
		}
	}

	return count, nil
}
