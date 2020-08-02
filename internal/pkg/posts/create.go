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

	return p.getByID(ctx, postID)
}

func (p *Posts) getByID(ctx context.Context, postID string) (models.Post, error) {
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.text, p.upvote as UpvotePercentage, p.created, " +
		"p.category, CASE WHEN p.is_link IS TRUE THEN 'link' ELSE 'text' END AS type " +
		"FROM posts p " +
		"WHERE p.id = $1"
	var post models.Post
	err := p.db.Sqlx.GetContext(ctx, &post, query, postID)
	return post, err
}
