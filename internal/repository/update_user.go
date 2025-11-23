package repository

import (
	"context"
	"fmt"
)

const qUpdateUser = `UPDATE users SET is_active = $1 WHERE user_id = $2`

func (r *RepositoryImpl) UpdateUser(ctx context.Context, userId string, isActive bool) error {
	_, err := r.Postgres.DB.Exec(ctx, qUpdateUser, isActive, userId)
	if err != nil {
		return fmt.Errorf("Ошибка в exec, метод UpdateUser: %w", err)
	}
	return nil
}
