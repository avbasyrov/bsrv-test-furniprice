package main

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/posts"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/routes"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/users"
	"log"
	"net/http"
)

func main() {
	authSecret := []byte("some secret key")
	ctx := context.Background()

	cfg := config.New()
	db := dbcon.New(ctx, cfg.DB)
	postsRepo := posts.New(db)
	usersRepo := users.New(db)

	httpHandler := routes.New(authSecret, usersRepo, postsRepo).InitRoutes()

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", httpHandler)
}
