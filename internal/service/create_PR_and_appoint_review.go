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

func (s *ServiceImpl) CreatePRAndAppointReview(ctx context.Context, createPR entity.CreatePRRequest) (entity.PullRequest, error) {
	user, err := s.Repo.GetUser(ctx, createPR.AuthorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.PullRequest{}, fmt.Errorf("%w", Error.NOT_FOUND)
		}
		s.Log.Error("Ошибка при получении юзера, метод CreatePRAndAppointReview", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при получении юзера, метод CreatePRAndAppointReview: %w", err)
	}

	_, err = s.Repo.GetPR(ctx, createPR.PullRequestID)
	if err == nil {
		return entity.PullRequest{}, fmt.Errorf("%w", Error.PR_EXISTS)
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		s.Log.Error("Ошибка при получении PR, метод CreatePRAndAppointReview", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при получении PR, метод CreatePRAndAppointReview: %w", err)
	}

	team, err := s.Repo.GetTeam(ctx, user.TeamName)
	if err != nil {
		s.Log.Error("Ошибка при получении команды, метод CreatePRAndAppointReview", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при получении команды, метод CreatePRAndAppointReview: %w", err)
	}

	users, err := s.Repo.GetUsersWithTeam(ctx, team.TeamName)
	if err != nil {
		s.Log.Error("Ошибка при получении юзеров одной команды, метод CreatePRAndAppointReview", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при получении юзеров одной команды, метод CreatePRAndAppointReview: %w", err)
	}

	AssignedReviewers := make([]string, 0, 2)
	for _, u := range users {
		if u.UserId != user.UserId && u.IsActive == true && len(AssignedReviewers) < 2 {
			AssignedReviewers = append(AssignedReviewers, u.UserId)
		}
	}

	PR := entity.PullRequest{
		PullRequestId:     createPR.PullRequestID,
		PullRequestName:   createPR.PullRequestName,
		AuthorId:          createPR.AuthorID,
		Status:            "Open",
		AssignedReviewers: AssignedReviewers,
	}

	if err := s.Repo.CreatePR(ctx, PR); err != nil {
		s.Log.Error("Ошибка при создании PR, метод CreatePRAndAppointReview", zap.Error(err))
		return entity.PullRequest{}, fmt.Errorf("Ошибка при создании PR, метод CreatePRAndAppointReview: %w", err)
	}

	return PR, nil
}
