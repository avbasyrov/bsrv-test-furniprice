package models

import (
	"context"
	"time"
)

type Post struct {
	Id               string    `json:"id"`
	Score            int       `json:"score"`
	Views            int       `json:"views"`
	Title            string    `json:"title"`
	Url              string    `json:"url"`
	Text             string    `json:"text"`
	UpvotePercentage int       `json:"upvotePercentage"`
	Created          time.Time `json:"created"`
	Category         string    `json:"category"`
	Type             string    `json:"type"`
}

type PostRepository interface {
	Create(ctx context.Context, title string, authorID int, url string, text string, category string, isLink bool) (Post, error)
	List(context.Context) ([]Post, error)
	ByCategory(ctx context.Context, category string) ([]Post, error)
}
