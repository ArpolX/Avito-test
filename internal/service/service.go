package service

import (
	"avito-test/internal/entity"
	"avito-test/internal/repository"
	"context"

	"go.uber.org/zap"
)

type Service interface {
	CreateTeamWithUsers(ctx context.Context, team entity.Team) (entity.Team, error)
	GetTeamWithUsers(ctx context.Context, teamName string) (entity.Team, error)
	SetFlagIsActive(ctx context.Context, setActive entity.SetIsActiveRequest) (entity.User, error)
	CreatePRAndAppointReview(ctx context.Context, createPR entity.CreatePRRequest) (entity.PullRequest, error)
	MarkPRMERGED(ctx context.Context, merge entity.MergePRRequest) (entity.PullRequest, error)
	RemapReview(ctx context.Context, remap entity.RemapReview) (entity.PullRequest, error)
	GetPRWithUserReview(ctx context.Context, userId string) ([]entity.PullRequestShort, error)
	AmountPROpen(ctx context.Context) (int, error)
	TeamSwitchActive(ctx context.Context, teamName string) (entity.Team, error)
}

type ServiceImpl struct {
	Log  *zap.Logger
	Repo repository.Repository
}

func NewServiceImpl(log *zap.Logger, repo repository.Repository) Service {
	return &ServiceImpl{
		Log:  log,
		Repo: repo,
	}
}
