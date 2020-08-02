package models

type Post struct {
	Id      string
	Score   int
	Views   int
	Title   string
	Url     string
	Upvote  int
	Created int // timestamp
}

type PostRepository interface {
	List() (error, []Post)
}
