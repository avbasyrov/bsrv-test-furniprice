package posts

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/models"
)

func (p *Posts) Create(ctx context.Context, title string, authorID int, url string, text string, category string, isLink bool) (models.Post, error) {
	const query = "INSERT INTO public.posts (created, title, author_id, url, text, category, is_link) " +
		"VALUES (NOW(), $1, $2, $3, $4, $5, $6) RETURNING id"

	postID := ""
	row := p.db.Sqlx.QueryRowContext(ctx, query, title, authorID, url, text, category, isLink)
	err := row.Scan(&postID)
	if err != nil {
		return models.Post{}, err
	}

	return p.GetByID(ctx, postID)
}
