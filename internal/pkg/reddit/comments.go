package reddit

import (
	"context"
)

func (r *Reddit) CreateComment(ctx context.Context, userID int, postID string, text string) error {
	return r.comments.Add(ctx, postID, text, userID)
}

func (r *Reddit) DeleteComment(ctx context.Context, userID int, commentID string) error {
	return r.comments.Delete(ctx, commentID, userID)
}
