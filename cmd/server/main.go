package main

import (
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/routes"
	"log"
	"net/http"
)

func main() {
	c := config.New()
	log.Println(c.DB.DbName + " at " + c.DB.User + ":" + c.DB.Password + "@" + c.DB.Host + ":" + c.DB.Port)
	r := routes.InitRoutes()

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}
