package controller

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func ValidJson(r *http.Request) *jsoniter.Decoder {
	decoder := Json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder
}
