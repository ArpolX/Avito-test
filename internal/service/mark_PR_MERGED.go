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

func (s *ServiceImpl) MarkPRMERGED(ctx context.Context, merge entity.MergePRRequest) (entity.PullRequest, error) {
	if err := s.Repo.TagMerge(ctx, merge.PullRequestID); err != nil {
		s.Log.Error("Ошибка при merge, метод MarkPRMERGED", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при merge, метод MarkPRMERGED: %w", err)
	}

	pr, err := s.Repo.GetPR(ctx, merge.PullRequestID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.PullRequest{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка при получении pr, метод MarkPRMERGED", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при получении pr, метод MarkPRMERGED: %w", err)
	}

	return pr, nil
}
