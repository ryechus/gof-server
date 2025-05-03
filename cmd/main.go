package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers"
)

func main() {
	// mux := http.NewServeMux()
	chi_r := chi.NewRouter()

	chi_r.Use(middleware.Logger)

	chi_r.Get("/ping", handlers.HandlePing)

	// flags
	chi_r.Get("/getFlag/{flagKey}", handlers.GetFlagWithVariations)
	chi_r.Post("/evaluateFlag/{flagKey}", handlers.EvaluateFlag)
	chi_r.Post("/createFlag", handlers.CreateFlag)
	chi_r.Put("/updateFlag", handlers.UpdateFlag)

	// rules
	chi_r.Put("/rule", handlers.PutRule)

	// variations
	chi_r.Get("/getVariations/{flagKey}", handlers.GetFlagVariations)

	log.Println("Server is running on http://localhost:23456")
	ctx := context.Background()
	server := &http.Server{
		Addr:    ":23456",
		Handler: chi_r,
		BaseContext: func(l net.Listener) context.Context {
			storageIface := config.FlagStorageIface
			ctx = context.WithValue(ctx, config.KeyVariable, storageIface)
			return ctx
		},
	}

	log.Fatal(server.ListenAndServe())
}
