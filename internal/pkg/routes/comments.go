package routes

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

func (c *Routes) addComment(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		jsonReply(w, http.StatusBadRequest, "unknown payload")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()

	fd := &struct {
		Comment string `json:"comment"`
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

	postID := chi.URLParam(r, "post_id")

	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "auth error")
		return
	}

	err = c.reddit.CreateComment(r.Context(), userID, postID, fd.Comment)

	if err != nil {
		jsonReply(w, http.StatusInternalServerError, "can't add comment")
		return
	}

	c.respondWithPost(w, r, postID)
}

func (c *Routes) deleteComment(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "post_id")
	commentID := chi.URLParam(r, "comment_id")

	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "auth error")
		return
	}

	err = c.reddit.DeleteComment(r.Context(), userID, commentID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, "can't remove comment")
		return
	}

	c.respondWithPost(w, r, postID)
}
