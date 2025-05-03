package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers"
)

func main() {
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

	var portNumber int
	var host string
	flag.IntVar(&portNumber, "port-number", 23456, "port number of server")
	flag.StringVar(&host, "host", "127.0.0.1", "host of the server")
	flag.Parse()

	log.Printf("Server is running on http://%s:%d", host, portNumber)
	ctx := context.Background()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, portNumber),
		Handler: chi_r,
		BaseContext: func(l net.Listener) context.Context {
			storageIface := config.FlagStorageIface
			ctx = context.WithValue(ctx, config.KeyVariable, storageIface)
			return ctx
		},
	}

	log.Fatal(server.ListenAndServe())
}
