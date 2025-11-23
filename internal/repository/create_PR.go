package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qCreatePR = `
		INSERT INTO pull_requests 
		(pull_request_id, pull_request_name, author_id, status)
		VALUES ($1, $2, $3, $4)
	`

const qReviewers = `
		INSERT INTO pull_request_reviewers (pr_id, reviewer_id)
		VALUES ($1, $2)
	`

func (r *RepositoryImpl) CreatePR(ctx context.Context, pr entity.PullRequest) error {
	_, err := r.Postgres.DB.Exec(ctx, qCreatePR, pr.PullRequestId, pr.PullRequestName, pr.AuthorId, pr.Status)
	if err != nil {
		return fmt.Errorf("Ошибка в Exec, метод CreatePR: %w", err)
	}

	for _, reviewer := range pr.AssignedReviewers {
		_, err := r.Postgres.DB.Exec(ctx, qReviewers, pr.PullRequestId, reviewer)
		if err != nil {
			return fmt.Errorf("Ошибка в Exec, метод CreatePR: %w", err)
		}
	}

	return nil
}
