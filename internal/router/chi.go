package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/placer14/gof-server/internal/handlers"
)

func GetChiRouter() http.Handler {
	chi_r := chi.NewRouter()

	chi_r.Use(middleware.Logger)
	chi_r.Use(middleware.Recoverer)

	chi_r.Use(middleware.Heartbeat("/ping"))
	// flags
	chi_r.Get("/getFlag/{flagKey}", handlers.GetFlagWithVariations)
	chi_r.Post("/evaluateFlag/{flagKey}", handlers.EvaluateFlag)
	chi_r.Post("/createFlag", handlers.CreateFlag)
	chi_r.Put("/updateFlag", handlers.UpdateFlag)

	// rules
	chi_r.Put("/rule", handlers.PutRule)

	// variations
	chi_r.Get("/getVariations/{flagKey}", handlers.GetFlagVariations)

	return chi_r
}
