package route

import (
	"avito-test/internal/controller"

	"github.com/go-chi/chi"
)

func Handlers(ctrl controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/team", func(r chi.Router) {
		r.Post("/add", ctrl.CreateTeamWithUsers)
		r.Get("/get", ctrl.GetTeamWithUsers)
		r.Post("/switchFalse", ctrl.TeamSwitchActive)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/setIsActive", ctrl.SetFlagIsActive)
		r.Get("/getReview", ctrl.GetPRWithUserReview)
	})

	r.Route("/pullRequest", func(r chi.Router) {
		r.Post("/create", ctrl.CreatePRAndAppointReview)
		r.Post("/merge", ctrl.MarkPRMERGED)
		r.Post("/reassign", ctrl.RemapReview)

	})

	r.Get("/amount/prOpen", ctrl.AmountPROpen)

	return r
}
