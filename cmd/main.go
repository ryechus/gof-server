package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/placer14/gof-server/internal/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.PingHandler)

	http.HandleFunc("/string/{flagKey}", handlers.StringFlagHandler)
	http.HandleFunc("/float/{flagKey}", handlers.FloatFlagHandler)
	http.HandleFunc("/int/{flagKey}", handlers.IntFlagHandler)
	http.HandleFunc("/bool/{flagKey}", handlers.BoolFlagHandler)

	http.HandleFunc("/setFlagValue/{flagKey}", handlers.SetFlagValue)

	fmt.Println("Server is running on http://localhost:23456")
	log.Fatal(http.ListenAndServe(":23456", nil))
}
