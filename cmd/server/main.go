package main

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/auth"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/posts"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/session"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/users"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/routes"
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
	sessionManager := session.New(db)
	authManager := auth.New(authSecret, sessionManager, usersRepo)

	httpHandler := routes.New(authSecret, authManager, usersRepo, postsRepo).
		InitRoutes()

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", httpHandler)
}
