package repository

import (
	"avito-test/internal/entity"
	"context"
	"fmt"
)

const qGetAllPR = `select status from pull_requests`

func (r *RepositoryImpl) GetAllPR(ctx context.Context) ([]entity.PullRequestShort, error) {
	rows, err := r.Postgres.DB.Query(ctx, qGetAllPR)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в exec, метод GetAllPR: %w", err)
	}
	defer rows.Close()

	var prs []entity.PullRequestShort

	for rows.Next() {
		var pr entity.PullRequestShort
		err := rows.Scan(
			&pr.Status,
		)
		if err != nil {
			return nil, err
		}
		prs = append(prs, pr)
	}

	return prs, nil
}
