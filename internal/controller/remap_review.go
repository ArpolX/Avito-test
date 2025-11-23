package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) RemapReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var re entity.RemapReview

	if err := Json.NewDecoder(r.Body).Decode(&re); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/reassign, метод RemapReview", zap.Error(err))
		CreateError("400", "Ошибка валидации запроса", w)
		return
	}

	PROrig, err := c.Srv.RemapReview(ctx, re)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		if errors.Is(err, Error.PR_MERGED) {
			CreateError("PR_MERGED", err.Error(), w)
			return
		}
		if errors.Is(err, Error.NOT_ASSIGNED) {
			CreateError("NOT_ASSIGNED", err.Error(), w)
			return
		}
		if errors.Is(err, Error.NO_CANDIDATE) {
			CreateError("NO_CANDIDATE", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /pullRequest/reassign, метод RemapReview", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&PROrig); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/reassign, метод RemapReview", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
