package repository

import (
	"avito-test/internal/entity"
	db "avito-test/internal/infrastructure/database/postgres"
	"context"

	"go.uber.org/zap"
)

type Repository interface {
	CreateTeamWithUsers(ctx context.Context, teamName string, users []entity.TeamMember) error
	GetTeam(ctx context.Context, teamName string) (entity.Team, error)
	UpdateUser(ctx context.Context, userId string, isActive bool) error
	GetUser(ctx context.Context, userId string) (entity.User, error)
	CreatePR(ctx context.Context, pr entity.PullRequest) error
	TagMerge(ctx context.Context, prId string) error
	UpdatePR(ctx context.Context, pr entity.PullRequest) error
	GetPR(ctx context.Context, prId string) (entity.PullRequest, error)
	GetUsersWithTeam(ctx context.Context, teamName string) ([]entity.User, error)
	GetPRSWithUser(ctx context.Context, userId string) ([]entity.PullRequestShort, error)
	GetAllPR(ctx context.Context) ([]entity.PullRequestShort, error)
}

type RepositoryImpl struct {
	Log      *zap.Logger
	Postgres *db.Postgres
}

func NewRepositoryImpl(postgres *db.Postgres, log *zap.Logger) Repository {
	return &RepositoryImpl{
		Log:      log,
		Postgres: postgres,
	}
}
