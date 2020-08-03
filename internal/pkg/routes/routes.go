package routes

import (
	"context"
	"encoding/json"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
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

type Routes struct {
	authSecret []byte
	auth       interfaces.AuthManager
	users      interfaces.UsersRepository
	posts      interfaces.PostRepository
	comments   interfaces.CommentsRepository
	reddit     interfaces.RedditService
}

func New(authSecret []byte,
	auth interfaces.AuthManager,
	users interfaces.UsersRepository,
	posts interfaces.PostRepository,
	comments interfaces.CommentsRepository,
	reddit interfaces.RedditService,
) *Routes {
	return &Routes{
		auth:       auth,
		authSecret: authSecret,
		users:      users,
		posts:      posts,
		comments:   comments,
		reddit:     reddit,
	}
}

func (c *Routes) InitRoutes() *chi.Mux {
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

	r.Use(commonHeaders)

	// Routes for index.html & /static/**
	staticRoutes(r)

	r.Post("/api/login", c.login)
	r.Post("/api/register", c.register)
	r.Post("/api/posts", c.createPost)
	r.Get("/api/post/{post_id}", c.getByID)
	r.Get("/api/posts/", c.listPosts)
	r.Get("/api/user/{user_name}", c.getByAuthor)
	r.Get("/api/post/{post_id}/upvote", c.upVote)
	r.Get("/api/post/{post_id}/unvote", c.unVote)
	r.Get("/api/post/{post_id}/downvote", c.downVote)
	r.Delete("/api/post/{post_id}", c.deletePost)
	r.Post("/api/post/{post_id}", c.addComment)
	r.Delete("/api/post/{post_id}/{comment_id}", c.deleteComment)

	r.Get("/api/posts/*", func(w http.ResponseWriter, r *http.Request) {
		myUrl, err := url.Parse(r.URL.Path)
		if err != nil {
			log.Fatal(err)
		}
		category := path.Base(myUrl.Path)

		ctx := context.Background()
		allPosts, err := c.posts.ByCategory(ctx, category)
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
		log.Println(status, string(jsonData))
		return jsonData, status
	}

	jsonData, err = json.Marshal(data)
	if err != nil {
		jsonData = []byte("Can't marshall to JSON: " + err.Error())
		status = 500
	}

	log.Println(status, string(jsonData))
	return jsonData, status
}

func staticRoutes(r chi.Router) {
	workDir, _ := os.Getwd()

	webDir := http.Dir(filepath.Join(workDir, "web"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, string(webDir)+"/index.html")
	})
	r.Get("/u/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, string(webDir)+"/index.html")
	})
	r.Get("/a/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, string(webDir)+"/index.html")
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(""))
	})
	r.Head("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(""))
	})
	r.Get("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(""))
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

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func jsonReply(w http.ResponseWriter, status int, msg string) {
	resp, _ := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	w.WriteHeader(status)
	_, _ = w.Write(resp)
}
