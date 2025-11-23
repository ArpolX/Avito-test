package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) CreateTeamWithUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var t entity.Team

	if err := decoder.Decode(&t); err != nil {
		c.Log.Error("Ошибка обработки пути /team/add, метод CreateTeamWithUsers", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	teamResp, err := c.Srv.CreateTeamWithUsers(ctx, t)
	if err != nil {
		if errors.Is(err, Error.TEAM_EXISTS) {
			CreateError("TEAM_EXISTS", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /team/add, метод CreateTeamWithUsers", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&teamResp); err != nil {
		c.Log.Error("Ошибка обработки пути /team/add, метод CreateTeamWithUsers", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
