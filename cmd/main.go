package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"os"

	"github.com/placer14/gof-server/internal/config"
	"github.com/placer14/gof-server/internal/router"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var portNumber int
	var host string
	flag.IntVar(&portNumber, "port-number", 23456, "port number of server")
	flag.StringVar(&host, "host", "127.0.0.1", "host of the server")
	flag.Parse()

	log.Info().Msgf("Server is running on http://%s:%d", host, portNumber)
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

	log.Fatal().Msg(server.ListenAndServe().Error())
}
