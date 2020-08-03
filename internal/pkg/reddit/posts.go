package reddit

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
)

func (r *Reddit) CreatePost(ctx context.Context, userID int, title string,
	url string, text string, category string, isLink bool,
) (interfaces.Post, error) {
	return r.posts.Create(ctx, title, userID, url, text, category, isLink)
}

func (r *Reddit) DeletePost(ctx context.Context, userID int, postID string) error {
	return r.posts.Delete(ctx, postID, userID)
}

func (r *Reddit) PostGetByID(ctx context.Context, postID string) (interfaces.Post, error) {
	return r.posts.GetByID(ctx, postID)
}

func (r *Reddit) PostsGetByAuthor(ctx context.Context, userID int) ([]interfaces.Post, error) {
	return r.posts.ByAuthor(ctx, userID)
}

func (r *Reddit) PostsGetByCategory(ctx context.Context, category string) ([]interfaces.Post, error) {
	return r.posts.ByCategory(ctx, category)
}

func (r *Reddit) PostsGetAll(ctx context.Context) ([]interfaces.Post, error) {
	return r.posts.List(ctx)
}
