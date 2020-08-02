package routes

import (
	"context"
	"encoding/json"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/posts"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func InitRoutes(postsRepo *posts.Posts) *chi.Mux {
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
		ctx := context.Background()
		allPosts, err := postsRepo.List(ctx)
		jsonData, status := toJSON(allPosts, err)
		log.Println(status, string(jsonData))

		w.WriteHeader(status)

		_, err = w.Write(jsonData)
		if err != nil {
			log.Println(err)
		}
	})

	r.Get("/api/posts/*", func(w http.ResponseWriter, r *http.Request) {
		myUrl, err := url.Parse(r.URL.Path)
		if err != nil {
			log.Fatal(err)
		}
		category := path.Base(myUrl.Path)

		ctx := context.Background()
		allPosts, err := postsRepo.ByCategory(ctx, category)
		jsonData, status := toJSON(allPosts, err)
		log.Println(status, string(jsonData))

		w.WriteHeader(status)

		_, err = w.Write(jsonData)
		if err != nil {
			log.Println(err)
		}
	})

	return r
}

func toJSON(data interface{}, err error) ([]byte, int) {
	var jsonData []byte
	status := 200 // HTTP: OK

	if err != nil {
		jsonData = []byte("Can't load data: " + err.Error())
		status = 500
		return jsonData, status
	}

	jsonData, err = json.Marshal(data)
	if err != nil {
		jsonData = []byte("Can't marshall to JSON: " + err.Error())
		status = 500
	}

	return jsonData, status
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
