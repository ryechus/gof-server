package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/router"
)

func main() {
	var portNumber int
	var host string
	flag.IntVar(&portNumber, "port-number", 23456, "port number of server")
	flag.StringVar(&host, "host", "127.0.0.1", "host of the server")
	flag.Parse()

	log.Printf("Server is running on http://%s:%d", host, portNumber)
	ctx := context.Background()
	chiRouter := router.GetChiRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, portNumber),
		Handler: chiRouter,
		BaseContext: func(l net.Listener) context.Context {
			storageIface := config.FlagStorageIface
			ctx = context.WithValue(ctx, config.KeyVariable, storageIface)
			return ctx
		},
	}

	log.Fatal(server.ListenAndServe())
}
