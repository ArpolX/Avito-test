package controller

import (
	"avito-test/internal/entity"
	Error "avito-test/pkg/errors"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (c *ControllerImpl) MarkPRMERGED(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := ValidJson(r)
	var m entity.MergePRRequest

	if err := decoder.Decode(&m); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/merge, метод MarkPRMERGED", zap.Error(err))
		CreateError("400", fmt.Sprintf("Ошибка валидации запроса, проверьте теги: %v", err), w)
		return
	}

	PROrig, err := c.Srv.MarkPRMERGED(ctx, m)
	if err != nil {
		if errors.Is(err, Error.NOT_FOUND) {
			CreateError("NOT_FOUND", err.Error(), w)
			return
		}
		c.Log.Error("Ошибка обработки пути /pullRequest/merge, метод MarkPRMERGED", zap.Error(err))
		CreateError("400", err.Error(), w)
		return
	}

	if err := Json.NewEncoder(w).Encode(&PROrig); err != nil {
		c.Log.Error("Ошибка обработки пути /pullRequest/merge, метод MarkPRMERGED", zap.Error(err))
		CreateError("400", err.Error(), w)
	}
}
