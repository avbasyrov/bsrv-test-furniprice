package interfaces

import "context"

type RedditService interface {
	// Comments
	CreateComment(ctx context.Context, userID int, postID string, text string) error
	DeleteComment(ctx context.Context, userID int, commentID string) error

	// Votes
	UpVote(ctx context.Context, userID int, postID string) error
	UnVote(ctx context.Context, userID int, postID string) error
	DownVote(ctx context.Context, userID int, postID string) error

	// Posts
	CreatePost(ctx context.Context, userID int, title string, url string, text string, category string, isLink bool) (Post, error)
	DeletePost(ctx context.Context, userID int, postID string) error
	PostGetByID(ctx context.Context, postID string) (Post, error)
	PostsGetByAuthor(ctx context.Context, userID int) ([]Post, error)
	PostsGetByCategory(ctx context.Context, category string) ([]Post, error)
	PostsGetAll(ctx context.Context) ([]Post, error)
}
