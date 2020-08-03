package routes

import (
	"encoding/json"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
	"log"
	"net/http"
)

func (c *Routes) respondWithPost(w http.ResponseWriter, r *http.Request, postID string) {
	postData, err := c.reddit.PostGetByID(r.Context(), postID)
	jsonData, status := toJSON(preparePostToJSON(postData), err)

	w.WriteHeader(status)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

func preparePostToJSON(p interfaces.Post) Post {
	votesJson := json.RawMessage(p.Votes)
	commentsJson := json.RawMessage(p.Comments)

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
		Votes:            &votesJson,
		Comments:         &commentsJson,
	}
}

func jsonReply(w http.ResponseWriter, status int, msg string) {
	resp, _ := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	w.WriteHeader(status)
	_, _ = w.Write(resp)
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
