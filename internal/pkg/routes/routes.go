package routes

import (
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/posts"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes for index.html & /static/**
	staticRoutes(r)

	r.Get("/api/posts/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(posts.List())
	})

	return r
}

func staticRoutes(r chi.Router) {
	workDir, _ := os.Getwd()

	webDir := http.Dir(filepath.Join(workDir, "web"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(string(webDir) + "/index.html")
		http.ServeFile(w, r, string(webDir)+"/index.html")
	})

	filesDir := http.Dir(filepath.Join(workDir, "web/static"))
	fileServer(r, "/static", filesDir)
}

// fileServer conveniently sets up a http.fileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
