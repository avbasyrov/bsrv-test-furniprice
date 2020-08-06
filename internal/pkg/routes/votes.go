package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (c *Routes) upVote(w http.ResponseWriter, r *http.Request) {
	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "unauthorized")
		return
	}

	postID := chi.URLParam(r, "post_id")

	err = c.reddit.UpVote(r.Context(), userID, postID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.respondWithPost(w, r, postID)
}

func (c *Routes) unVote(w http.ResponseWriter, r *http.Request) {
	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "unauthorized")
		return
	}

	postID := chi.URLParam(r, "post_id")

	err = c.reddit.UnVote(r.Context(), userID, postID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.respondWithPost(w, r, postID)
}

func (c *Routes) downVote(w http.ResponseWriter, r *http.Request) {
	userID, _, err := c.auth.GetAuthorized(r.Context(), r.Header.Get("authorization"))
	if err != nil {
		jsonReply(w, http.StatusForbidden, "unauthorized")
		return
	}

	postID := chi.URLParam(r, "post_id")

	err = c.reddit.DownVote(r.Context(), userID, postID)
	if err != nil {
		jsonReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.respondWithPost(w, r, postID)
}
