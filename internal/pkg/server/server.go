package server

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/auth"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/comments"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/posts"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/session"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dal/users"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/reddit"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/routes"
	"github.com/go-chi/chi"
)

type App struct {
	HttpHandler *chi.Mux
}

func New(authSecret []byte, cfg *config.Config) *App {
	db := dbcon.New(context.Background(), cfg.DB)

	postsRepo := posts.New(db)
	usersRepo := users.New(db)
	commentsRepo := comments.New(db)
	sessionManager := session.New(db)
	authManager := auth.New(authSecret, sessionManager, usersRepo)
	redditService := reddit.New(postsRepo, commentsRepo)

	httpHandler := routes.New(authSecret, authManager, usersRepo, redditService).
		InitRoutes()

	return &App{HttpHandler: httpHandler}
}
