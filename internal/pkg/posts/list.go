package posts

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/models"
)

type Posts struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Posts {
	return &Posts{
		db: db,
	}
}

func (p *Posts) List(ctx context.Context) ([]models.Post, error) {
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.upvote as UpvotePercentage, p.created, " +
		"c.title as category " +
		"FROM posts p " +
		"LEFT JOIN categories c ON c.id = p.category_id"
	var posts []models.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query)
	return posts, err
}

func (p *Posts) ByCategory(ctx context.Context, category string) ([]models.Post, error) {
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.upvote as UpvotePercentage, p.created, " +
		"c.title as category " +
		"FROM posts p " +
		"LEFT JOIN categories c ON c.id = p.category_id " +
		"WHERE c.title = $1"
	var posts []models.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query, category)
	return posts, err
}
