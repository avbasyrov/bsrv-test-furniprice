package routes

import (
	"log"
	"net/http"
)

func (c *Routes) listPosts(w http.ResponseWriter, r *http.Request) {
	allPosts, err := c.posts.List(r.Context())
	jsonData, status := toJSON(allPosts, err)
	log.Println(status, string(jsonData))

	w.WriteHeader(status)

	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}
