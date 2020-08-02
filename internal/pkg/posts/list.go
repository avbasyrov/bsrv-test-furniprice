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
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.text, p.upvote as UpvotePercentage, p.created, " +
		"p.category, p.author_id AS AuthorID, u.login AS AuthorName, " +
		"CASE WHEN p.is_link IS TRUE THEN 'link' ELSE 'text' END AS type  " +
		"FROM posts p " +
		"LEFT JOIN users u ON u.id = p.author_id "
	var posts []models.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query)
	return posts, err
}

func (p *Posts) ByCategory(ctx context.Context, category string) ([]models.Post, error) {
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.text, p.upvote as UpvotePercentage, p.created, " +
		"p.category, p.author_id AS AuthorID, u.login AS AuthorName, " +
		"CASE WHEN p.is_link IS TRUE THEN 'link' ELSE 'text' END AS type " +
		"FROM posts p " +
		"LEFT JOIN users u ON u.id = p.author_id " +
		"WHERE p.category = $1"
	var posts []models.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query, category)
	return posts, err
}

func (p *Posts) GetByID(ctx context.Context, postID string) (models.Post, error) {
	const query = "SELECT p.id, p.score, p.views, p.title, p.url, p.text, p.upvote as UpvotePercentage, p.created, " +
		"p.category, p.author_id AS AuthorID, u.login AS AuthorName, " +
		"CASE WHEN p.is_link IS TRUE THEN 'link' ELSE 'text' END AS type " +
		"FROM posts p " +
		"LEFT JOIN users u ON u.id = p.author_id " +
		"WHERE p.id = $1"
	var post models.Post
	err := p.db.Sqlx.GetContext(ctx, &post, query, postID)
	return post, err
}
