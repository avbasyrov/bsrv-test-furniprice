package routes

import (
	"encoding/json"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/models"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Id     string `json:"id"`
	Score  int    `json:"score"`
	Author struct {
		Username string `json:"username"`
		ID       int    `json:"id"`
	} `json:"author"`
	Views            int       `json:"views"`
	Title            string    `json:"title"`
	Url              string    `json:"url"`
	Text             string    `json:"text"`
	UpvotePercentage int       `json:"upvotePercentage"`
	Created          time.Time `json:"created"`
	Category         string    `json:"category"`
	Type             string    `json:"type"`
	// TODO:
	Votes    []string `json:"votes"`
	Comments []string `json:"comments"`
}

func (c *Routes) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if r.Header.Get("Content-Type") != "application/json" {
		jsonError(w, http.StatusBadRequest, "unknown payload")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()

	fd := &struct {
		Category string `json:"category"`
		Type     string `json:"type"`
		Url      string `json:"url"`
		Title    string `json:"title"`
		Text     string `json:"text"`
	}{}
	err := json.Unmarshal(body, fd)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "cant unpack payload")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		jsonError(w, http.StatusBadRequest, "unknown payload")
		return
	}

	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonError(w, http.StatusForbidden, "auth error")
		return
	}

	post, err := c.posts.Create(r.Context(), fd.Title, userID, fd.Url, fd.Text, fd.Category, fd.Type == "link")

	jsonData, status := toJSON(c.preparePostToJSON(post), err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) listPosts(w http.ResponseWriter, r *http.Request) {
	allPosts, err := c.posts.List(r.Context())
	jsonPosts := make([]Post, 0, len(allPosts))

	for _, postData := range allPosts {
		jsonPosts = append(jsonPosts, c.preparePostToJSON(postData))
	}

	jsonData, status := toJSON(jsonPosts, err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) getByID(w http.ResponseWriter, r *http.Request) {
	var jsonData []byte
	var status int

	postID := chi.URLParam(r, "post_id")
	postData, err := c.posts.GetByID(r.Context(), postID)
	if err != nil {
		jsonData, status = toJSON(nil, err)
	} else {
		jsonData, status = toJSON(c.preparePostToJSON(postData), err)
	}

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) preparePostToJSON(p models.Post) Post {
	return Post{
		Id:    p.Id,
		Score: p.Score,
		Author: struct {
			Username string `json:"username"`
			ID       int    `json:"id"`
		}{p.AuthorName, p.AuthorID},
		Views:            p.Views,
		Title:            p.Title,
		Url:              p.Url,
		Text:             p.Text,
		UpvotePercentage: p.UpvotePercentage,
		Created:          p.Created,
		Category:         p.Category,
		Type:             p.Type,
		Votes:            nil,
		Comments:         nil,
	}
}
