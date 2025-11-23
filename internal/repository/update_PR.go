package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qUpdatePR = `
		UPDATE pull_requests 
		SET pull_request_name = $1, status = $2
		WHERE pull_request_id = $3
	`

const qDeleteReviewers = `DELETE FROM pull_request_reviewers WHERE pr_id = $1`

func (r *RepositoryImpl) UpdatePR(ctx context.Context, pr entity.PullRequest) error {
	_, err := r.Postgres.DB.Exec(ctx, qUpdatePR, pr.PullRequestName, pr.Status, pr.PullRequestId)
	if err != nil {
		return fmt.Errorf("Ошибка в exec, метод UpdatePR: %w", err)
	}

	_, _ = r.Postgres.DB.Exec(ctx, qDeleteReviewers, pr.PullRequestId)

	for _, reviewer := range pr.AssignedReviewers {
		_, err := r.Postgres.DB.Exec(ctx, qReviewers, pr.PullRequestId, reviewer)
		if err != nil {
			return fmt.Errorf("Ошибка в exec, метод UpdatePR: %w", err)
		}
	}

	return nil
}
