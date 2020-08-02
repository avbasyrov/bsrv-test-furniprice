package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

	jsonData, status := toJSON(post, err)
	log.Println(status, string(jsonData))

	w.WriteHeader(status)

	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}
