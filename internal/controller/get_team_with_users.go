package controller

import (
	Error "avito-test/pkg/errors"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) GetTeamWithUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	teamName := r.URL.Query().Get("team_name")

	teamOrig, err := c.Srv.GetTeamWithUsers(ctx, teamName)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /team/get, метод GetTeamWithUsers", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&teamOrig); err != nil {
		c.Log.Error("Ошибка обработки пути /team/get, метод GetTeamWithUsers", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
