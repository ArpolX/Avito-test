package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qGetPR = `
		SELECT pull_request_id, pull_request_name, author_id, status, created_at, merged_at
		FROM pull_requests WHERE pull_request_id = $1
	`

const qGetReviewers = `SELECT reviewer_id FROM pull_request_reviewers WHERE pr_id = $1`

func (r *RepositoryImpl) GetPR(ctx context.Context, prId string) (entity.PullRequest, error) {

	var pr entity.PullRequest
	err := r.Postgres.DB.QueryRow(ctx, qGetPR, prId).Scan(
		&pr.PullRequestId,
		&pr.PullRequestName,
		&pr.AuthorId,
		&pr.Status,
		&pr.CreatedAt,
		&pr.MergedAt,
	)
	if err != nil {
		return entity.PullRequest{}, fmt.Errorf("Ошибка в QueryRow, метод GetPR: %w", err)
	}

	rows, err := r.Postgres.DB.Query(ctx, qGetReviewers, prId)
	if err != nil {
		return pr, fmt.Errorf("Ошибка в Query, метод GetPR: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rid string
		if err := rows.Scan(&rid); err != nil {
			return pr, err
		}
		pr.AssignedReviewers = append(pr.AssignedReviewers, rid)
	}

	return pr, nil
}
