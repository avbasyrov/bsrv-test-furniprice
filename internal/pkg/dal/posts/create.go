package posts

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
)

func (p *Posts) Create(ctx context.Context, title string, authorID int, url string, text string, category string, isLink bool) (interfaces.Post, error) {
	const query = "INSERT INTO public.posts (created, title, author_id, url, text, category, is_link) " +
		"VALUES (NOW(), $1, $2, $3, $4, $5, $6) RETURNING id"

	postID := ""
	row := p.db.Sqlx.QueryRowContext(ctx, query, title, authorID, url, text, category, isLink)
	err := row.Scan(&postID)
	if err != nil {
		return interfaces.Post{}, err
	}

	return p.GetByID(ctx, postID)
}

func (p *Posts) Delete(ctx context.Context, postID string, userID int) error {
	_, err := p.db.Sqlx.ExecContext(ctx, "DELETE FROM public.posts WHERE id = $1 AND author_id = $2", postID, userID)
	return err
}

func (p *Posts) IncrementViews(ctx context.Context, postID string) error {
	_, err := p.db.Sqlx.ExecContext(ctx, "UPDATE public.posts SET views = views + 1 WHERE id = $1", postID)
	return err
}
