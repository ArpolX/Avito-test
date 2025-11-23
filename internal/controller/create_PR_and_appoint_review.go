package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) CreatePRAndAppointReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var pr entity.CreatePRRequest

	if err := Json.NewDecoder(r.Body).Decode(&pr); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/create, метод CreatePRAndAppointReview", zap.Error(err))
		CreateError("400", "Ошибка валидации запроса", w)
		return
	}

	PR, err := c.Srv.CreatePRAndAppointReview(ctx, pr)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		if errors.Is(err, Error.PR_EXISTS) {
			CreateError("PR_EXISTS", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /pullRequest/create, метод CreatePRAndAppointReview", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&PR); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/create, метод CreatePRAndAppointReview", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
