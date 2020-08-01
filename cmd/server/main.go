package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"strings"
	"time"

	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
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
	RouteStatic(r)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}

func RouteStatic(r chi.Router) {
	workDir, _ := os.Getwd()

	webDir := http.Dir(filepath.Join(workDir, "web"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(string(webDir) + "/index.html")
		http.ServeFile(w, r, string(webDir)+"/index.html")
	})

	filesDir := http.Dir(filepath.Join(workDir, "web/static"))
	FileServer(r, "/static", filesDir)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
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
