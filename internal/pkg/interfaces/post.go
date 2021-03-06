package interfaces

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
	Votes            string
	Comments         string
}

type PostRepository interface {
	Create(ctx context.Context, title string, authorID int, url string, text string, category string, isLink bool) (Post, error)
	Delete(ctx context.Context, postID string, userID int) error
	List(ctx context.Context) ([]Post, error)
	ByCategory(ctx context.Context, category string) ([]Post, error)
	ByAuthor(ctx context.Context, authorID int) ([]Post, error)
	GetByID(ctx context.Context, postID string) (Post, error)
	IncrementViews(ctx context.Context, postID string) error
	VoteUp(ctx context.Context, postID string, userID int) error
	UnVote(ctx context.Context, postID string, userID int) error
	VoteDown(ctx context.Context, postID string, userID int) error
}
