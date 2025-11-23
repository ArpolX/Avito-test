package controller

import (
	"avito-test/internal/config"
	"avito-test/internal/service"
	"net/http"

	jsonIterator "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var Json = jsonIterator.ConfigCompatibleWithStandardLibrary

type Controller interface {
	CreateTeamWithUsers(w http.ResponseWriter, r *http.Request) // Team
	GetTeamWithUsers(w http.ResponseWriter, r *http.Request)    // Team
	SetFlagIsActive(w http.ResponseWriter, r *http.Request)
	CreatePRAndAppointReview(w http.ResponseWriter, r *http.Request)
	MarkPRMERGED(w http.ResponseWriter, r *http.Request)
	RemapReview(w http.ResponseWriter, r *http.Request)
	GetPRWithUserReview(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	Log *zap.Logger
	Srv service.Service
}

func NewControllerImpl(cfg config.Config, log *zap.Logger, srv service.Service) ControllerImpl {
	return ControllerImpl{
		Log: log,
		Srv: srv,
	}
}
