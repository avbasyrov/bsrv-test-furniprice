package models

import (
	"context"
	"time"
)

type Post struct {
	Id               string
	Score            int
	AuthorID         int
	AuthorName       string
	Views            int
	Title            string
	Url              string
	Text             string
	UpvotePercentage int
	Created          time.Time
	Category         string
	Type             string
}

type PostRepository interface {
	Create(ctx context.Context, title string, authorID int, url string, text string, category string, isLink bool) (Post, error)
	List(context.Context) ([]Post, error)
	ByCategory(ctx context.Context, category string) ([]Post, error)
	GetByID(ctx context.Context, postID string) (Post, error)
}
