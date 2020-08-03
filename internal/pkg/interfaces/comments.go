package interfaces

import (
	"context"
	"time"
)

type Comment struct {
	ID       string
	Body     string
	Created  time.Time
	AuthorID int
}

type CommentsRepository interface {
	Add(ctx context.Context, postID string, comment string, authorID int) error
	Delete(ctx context.Context, commentID string, authorID int) error
}
