package interfaces

import "context"

type RedditService interface {
	CreateComment(ctx context.Context, userID int, postID string, text string) error
	DeleteComment(ctx context.Context, userID int, commentID string) error
}
