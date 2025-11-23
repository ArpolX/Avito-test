package controller

import (
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) GetPRWithUserReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("user_id")

	PROrig, err := c.Srv.GetPRWithUserReview(ctx, userId)
	if err != nil {
		c.Log.Error("Ошибка обработки пути /users/getReview, метод GetPRWithUserReview", zap.Error(err))
		CreateError("400", "Ошибка валидации запроса", w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&PROrig); err != nil {
		c.Log.Error("Ошибка обработки пути /users/getReview, метод GetPRWithUserReview", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
