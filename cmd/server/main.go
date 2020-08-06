package main

import (
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/server"
	"log"
	"net/http"
)

func main() {
	authSecret := []byte("some secret key")
	cfg := config.New()

	app := server.New(authSecret, cfg)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", app.HttpHandler)
	if err != nil {
		log.Println(err)
	}
}
