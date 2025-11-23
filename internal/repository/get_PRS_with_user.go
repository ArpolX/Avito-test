package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qGetPRSWithUser = `
		SELECT pull_request_id, pull_request_name, author_id, status
		FROM pull_requests p
		JOIN pull_request_reviewers r ON r.pr_id = p.pull_request_id
		WHERE r.reviewer_id = $1
	`

func (r *RepositoryImpl) GetPRSWithUser(ctx context.Context, userId string) ([]entity.PullRequestShort, error) {
	rows, err := r.Postgres.DB.Query(ctx, qGetPRSWithUser, userId)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в exec, метод GetPRSWithUser: %w", err)
	}
	defer rows.Close()

	var prs []entity.PullRequestShort

	for rows.Next() {
		var pr entity.PullRequestShort
		err := rows.Scan(
			&pr.PullRequestId,
			&pr.PullRequestName,
			&pr.AuthorId,
			&pr.Status,
		)
		if err != nil {
			return nil, err
		}
		prs = append(prs, pr)
	}

	return prs, nil
}
