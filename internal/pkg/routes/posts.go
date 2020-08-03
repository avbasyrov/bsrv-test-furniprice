package routes

import (
	"encoding/json"
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
	Views            int              `json:"views"`
	Title            string           `json:"title"`
	Url              string           `json:"url"`
	Text             string           `json:"text"`
	UpvotePercentage int              `json:"upvotePercentage"`
	Created          time.Time        `json:"created"`
	Category         string           `json:"category"`
	Type             string           `json:"type"`
	Votes            *json.RawMessage `json:"votes"`
	Comments         *json.RawMessage `json:"comments"`
}

func (c *Routes) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		jsonReply(w, http.StatusBadRequest, "unknown payload")
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
		jsonReply(w, http.StatusBadRequest, "cant unpack payload")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		jsonReply(w, http.StatusBadRequest, "unknown payload")
		return
	}

	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "auth error")
		return
	}

	post, err := c.reddit.CreatePost(r.Context(), userID, fd.Title, fd.Url, fd.Text, fd.Category, fd.Type == "link")

	jsonData, status := toJSON(preparePostToJSON(post), err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) listPosts(w http.ResponseWriter, r *http.Request) {
	allPosts, err := c.reddit.PostsGetAll(r.Context())
	jsonPosts := make([]Post, 0, len(allPosts))

	for _, postData := range allPosts {
		jsonPosts = append(jsonPosts, preparePostToJSON(postData))
	}

	jsonData, status := toJSON(jsonPosts, err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) listPostsByCategory(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	allPosts, err := c.reddit.PostsGetByCategory(r.Context(), category)

	jsonPosts := make([]Post, 0, len(allPosts))
	for _, postData := range allPosts {
		jsonPosts = append(jsonPosts, preparePostToJSON(postData))
	}

	jsonData, status := toJSON(jsonPosts, err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) getByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")
	err := c.reddit.IncrementViews(r.Context(), postID)
	if err != nil {
		log.Println(err)
	}
	c.respondWithPost(w, r, postID)
}

func (c *Routes) getByAuthor(w http.ResponseWriter, r *http.Request) {
	authorName := chi.URLParam(r, "user_name")
	user, err := c.users.GetByLogin(r.Context(), authorName)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, "can't load given user")
		return
	}

	posts, err := c.reddit.PostsGetByAuthor(r.Context(), user.ID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, "can't load posts for given user")
		return
	}

	jsonPosts := make([]Post, 0, len(posts))

	for _, postData := range posts {
		jsonPosts = append(jsonPosts, preparePostToJSON(postData))
	}

	jsonData, status := toJSON(jsonPosts, err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func (c *Routes) deletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")

	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "auth error")
		return
	}

	err = c.reddit.DeletePost(r.Context(), userID, postID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, "can't remove POST")
	} else {
		jsonReply(w, http.StatusOK, "success")
	}
}
