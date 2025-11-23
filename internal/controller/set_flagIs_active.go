package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) SetFlagIsActive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var setActive entity.SetIsActiveRequest

	if err := decoder.Decode(&setActive); err != nil {
		c.Log.Error("Ошибка обработки пути /users/setIsActive, метод SetFlagIsActive", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	user, err := c.Srv.SetFlagIsActive(ctx, setActive)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /users/setIsActive, метод SetFlagIsActive", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&user); err != nil {
		c.Log.Error("Ошибка обработки пути /users/setIsActive, метод SetFlagIsActive", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
