package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.HandlePing)

	mux.HandleFunc("POST /evaluateFlag/{flagKey}", handlers.EvaluateFlag)

	mux.HandleFunc("POST /createFlag", handlers.CreateFlag)
	mux.HandleFunc("PUT /updateFlag", handlers.UpdateFlag)

	mux.HandleFunc("PUT /rule", handlers.PutRule)

	log.Println("Server is running on http://localhost:23456")
	ctx := context.Background()
	server := &http.Server{
		Addr:    ":23456",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			storageIface := config.FlagStorageIface
			ctx = context.WithValue(ctx, config.KeyVariable, storageIface)
			return ctx
		},
	}

	log.Fatal(server.ListenAndServe())
}
