package service

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (s *ServiceImpl) RemapReview(ctx context.Context, remap entity.RemapReview) (entity.PullRequest, error) {
	pr, err := s.Repo.GetPR(ctx, remap.PullRequestId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.PullRequest{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка GetPR", zap.Error(err))
		return entity.PullRequest{}, err
	}

	if pr.Status == "MERGED" {
		return entity.PullRequest{}, fmt.Errorf("%w", Error.PR_MERGED)
	}

	var found bool
	for _, r := range pr.AssignedReviewers {
		if r == remap.OldUserId {
			found = true
			break
		}
	}
	if !found {
		return entity.PullRequest{}, fmt.Errorf("%w", Error.NOT_ASSIGNED)
	}

	oldUser, err := s.Repo.GetUser(ctx, remap.OldUserId)
	if err != nil {
		return entity.PullRequest{}, fmt.Errorf("%w", Error.NOT_FOUND)
	}

	teamUsers, err := s.Repo.GetUsersWithTeam(ctx, oldUser.TeamName)
	if err != nil {
		s.Log.Error("Ошибка при получении users, метод RemapReview", zap.Error(err))
		return entity.PullRequest{}, err
	}

	// от дубликатов ревьюеров в рамках одного pr
	current := map[string]struct{}{}
	for _, r := range pr.AssignedReviewers {
		current[r] = struct{}{}
	}

	newReviewer := make([]string, 0, len(teamUsers)-2) // минус автор и действующий ревьюер
	for _, u := range teamUsers {
		if u.UserId != remap.OldUserId && u.UserId != pr.AuthorId && u.IsActive {
			if _, exists := current[u.UserId]; exists {
				continue
			}
			newReviewer = append(newReviewer, u.UserId)
		}
	}

	if len(newReviewer) == 0 {
		return entity.PullRequest{}, fmt.Errorf("%w", Error.NO_CANDIDATE)
	}

	for i, r := range pr.AssignedReviewers {
		if r == remap.OldUserId {
			pr.AssignedReviewers[i] = newReviewer[rand.Intn(len(newReviewer))]
			break
		}
	}

	if err := s.Repo.UpdatePR(ctx, pr); err != nil {
		s.Log.Error("Ошибка при обновлении pr, метод RemapReview", zap.Error(err))
		return entity.PullRequest{}, err
	}

	return pr, nil
}
