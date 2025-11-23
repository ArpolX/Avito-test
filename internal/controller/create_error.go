package controller

import (
	"avito-test/internal/entity"
	"encoding/json"
	"net/http"
)

func CreateError(code, message string, w http.ResponseWriter) {
	e := entity.ErrorResponse{
		Code:    code,
		Message: message,
	}

	json.NewEncoder(w).Encode(e)
}
