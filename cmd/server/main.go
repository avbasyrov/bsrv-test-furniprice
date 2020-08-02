package main

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/posts"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/routes"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	c := config.New()
	db := dbcon.New(ctx, c.DB)
	postsRepo := posts.New(db)

	r := routes.InitRoutes(postsRepo)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}
