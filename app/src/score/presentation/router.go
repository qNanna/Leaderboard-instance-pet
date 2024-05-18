package presentation

import (
	"net/http"

	middleware "main/app/src/middleware"

	chi "github.com/go-chi/chi/v5"
)

var scoreController = &ScoreController{}

func Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", scoreController.GetScore)
	router.With(middleware.TokenGuard).Post("/", scoreController.SaveScore)

	return router
}
