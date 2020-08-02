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
	UpvotePercentage int       `json:"upvotePercentage"`
	Created          time.Time `json:"created"`
	Category         string    `json:"category"`
}

type PostRepository interface {
	List(context.Context) ([]Post, error)
	ByCategory(ctx context.Context, category string) ([]Post, error)
}
