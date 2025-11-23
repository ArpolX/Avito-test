package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) TeamSwitchActive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var team entity.TeamSwitchActiveRequest

	if err := decoder.Decode(&team); err != nil {
		c.Log.Error("Ошибка обработки пути /team/switchFalse, метод TeamSwitchActive", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	teamOrig, err := c.Srv.TeamSwitchActive(ctx, team.TeamName)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /team/switchFalse, метод TeamSwitchActive", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&teamOrig); err != nil {
		c.Log.Error("Ошибка обработки пути /team/switchFalse, метод TeamSwitchActive", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
