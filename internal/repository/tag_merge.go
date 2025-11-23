package repository

import (
	"context"
	"fmt"
)

const qTagMerge = `
		UPDATE pull_requests 
		SET status = 'MERGED', merged_at = NOW()
		WHERE pull_request_id = $1
	`

func (r *RepositoryImpl) TagMerge(ctx context.Context, prId string) error {
	_, err := r.Postgres.DB.Exec(ctx, qTagMerge, prId)
	if err != nil {
		return fmt.Errorf("Ошибка в exec, метод TagMerge: %w", err)
	}
	return nil
}
