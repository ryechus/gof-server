package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/placer14/gof-server/internal/handlers"
	"github.com/placer14/gof-server/internal/provider"
)

func main() {
	mux := http.NewServeMux()
	provider.PopulateFlagValues()
	// db := database.GetDB()
	mux.HandleFunc("/ping", handlers.HandlePing)
	mux.HandleFunc("GET /string/{flagKey}", handlers.GetStringValue)
	mux.HandleFunc("POST /string/{flagKey}", handlers.SetStringvalue)
	mux.HandleFunc("GET /float/{flagKey}", handlers.GetFloatValue)
	mux.HandleFunc("GET /int/{flagKey}", handlers.GetIntValue)
	mux.HandleFunc("GET /bool/{flagKey}", handlers.GetBoolValue)

	fmt.Println("Server is running on http://localhost:23456")
	// defer db.Close()
	ctx := context.Background()
	server := &http.Server{
		Addr:    ":23456",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			m := provider.NewProviderMock()
			ctx = context.WithValue(ctx, provider.KeyFlagStore, m)
			return ctx
		},
	}

	log.Fatal(server.ListenAndServe())
}
