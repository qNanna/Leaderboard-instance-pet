package router

import (
	middleware "main/app/src/middleware"
	score "main/app/src/score/presentation"

	chi "github.com/go-chi/chi/v5"
	chiMid "github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logging)
	router.Use(chiMid.SetHeader("Content-Type", "application/json"))

	router.Mount("/score", score.Router())

	return router
}
